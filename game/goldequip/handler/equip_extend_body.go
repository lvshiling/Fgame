package handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	//processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_EXTEND_BODY_TYPE), dispatch.HandlerFunc(handleGoldEquipExtendBody))
}

//处理金装继承信息
func handleGoldEquipExtendBody(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理金装继承信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSGoldEquipExtendBody)
	slotId := csMsg.GetSlotId()
	useItemIndex := csMsg.GetUseItemIndex()

	posType := inventorytypes.BodyPositionType(slotId)
	if !posType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"slotId":   slotId,
				"error":    err,
			}).Warn("goldequip:处理装备槽金装继承,位置错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = goldequipExtend(tpl, posType, useItemIndex)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotId":       slotId,
				"useItemIndex": useItemIndex,
				"error":        err,
			}).Error("goldequip:处理金装继承信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"slotId":       slotId,
			"useItemIndex": useItemIndex,
		}).Debug("goldequip:处理金装继承完成")
	return nil
}

//金装继承的逻辑
func goldequipExtend(pl player.Player, posType inventorytypes.BodyPositionType, useItemIndex int32) (err error) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipBag := goldequipManager.GetGoldEquipBag()
	targetIt := equipBag.GetByPosition(posType)
	useIt := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, useItemIndex)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"posType":      posType,
				"useItemIndex": useItemIndex,
			}).Warn("goldequip:继承失败,强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	//材料不存在
	if useIt == nil || useIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"posType":      posType,
				"useItemIndex": useItemIndex,
			}).Warn("goldequip:继承失败,强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	useItemTemplate := item.GetItemService().GetItem(int(useIt.ItemId))
	if !useItemTemplate.IsGoldEquip() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"posType":      posType,
				"useItemIndex": useItemIndex,
				"itemId":       useIt.ItemId,
			}).Warn("goldequip:继承失败,目标不是元神金装")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	targetPropertyData := targetIt.GetPropertyData().(*goldequiptypes.GoldEquipPropertyData)
	targetUpstar := targetPropertyData.UpstarLevel
	usePropertyData := useIt.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	useUpstar := usePropertyData.UpstarLevel
	if useUpstar <= targetUpstar {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"posType":      posType,
				"targetUpstar": targetUpstar,
				"useUpstar":    useUpstar,
			}).Warn("goldequip:金装继承失败，该装备强化等级低于当前装备，无法进行继承")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipCantNotUseExtend)
		return
	}

	jichengTemp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipJiChengTemplate(useUpstar)
	if jichengTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"posType":   posType,
				"useUpstar": useUpstar,
			}).Warn("goldequip:金装继承失败，继承配置模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//继承需要物品
	needItemId := jichengTemp.NeedItemId
	needItemCount := jichengTemp.NeedItemCount
	if !inventoryManager.HasEnoughItem(needItemId, needItemCount) {
		log.WithFields(
			log.Fields{
				"playerid":      pl.GetId(),
				"posType":       posType,
				"useIndex":      useItemIndex,
				"needItemId":    needItemId,
				"needItemCount": needItemCount,
			}).Warn("goldequip:金装继承失败,道具不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//消耗物品
	reason := commonlog.InventoryLogReasonGoldEquipExtendUse
	reasonText := fmt.Sprintf(reason.String(), useUpstar)
	flag := inventoryManager.UseItem(needItemId, needItemCount, reason, reasonText)
	if !flag {
		panic(fmt.Errorf("goldequip: goldequipExtend use item should be ok"))
	}

	//金装继承判断
	flag = equipBag.ExtendGoldEquipLevel(posType, useUpstar)
	if !flag {
		panic(fmt.Errorf("goldequip: goldequipExtend should be ok"))
	}
	inventoryManager.GoldEquipUpstarReturn(useItemIndex, 0)

	//金装继承日志
	extendReason := commonlog.GoldEquipLogReasonExtend
	extendReasonText := fmt.Sprintf(extendReason.String(), targetIt.GetItemId(), useIt.ItemId)
	eventData := goldequipeventtypes.CreatePlayerGoldEquipExtendLogEventData(targetUpstar, useUpstar, extendReason, extendReasonText)
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipExtendLog, pl, eventData)

	//同步属性
	goldequiplogic.GoldEquipPropertyChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGoldEquipExtendBody(int32(posType), useItemIndex, useUpstar)
	pl.SendMsg(scMsg)
	return
}
