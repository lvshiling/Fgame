package handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/goldequip/pbutil"
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
	//processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_EXTEND_BAG_TYPE), dispatch.HandlerFunc(handleGoldEquipExtendBag))
}

//处理背包金装继承信息
func handleGoldEquipExtendBag(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理背包金装继承信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSGoldEquipExtendBag)
	itemIndex := csMsg.GetItemIndex()
	useItemIndex := csMsg.GetUseItemIndex()

	err = goldequipExtendBag(tpl, itemIndex, useItemIndex)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"itemIndex":    itemIndex,
				"useItemIndex": useItemIndex,
				"error":        err,
			}).Error("goldequip:处理背包金装继承信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"itemIndex":    itemIndex,
			"useItemIndex": useItemIndex,
		}).Debug("goldequip:处理背包金装继承完成")
	return nil
}

//背包金装继承的逻辑
func goldequipExtendBag(pl player.Player, targetIndex, useIndex int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	targetIt := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, targetIndex)
	useIt := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, useIndex)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:继承失败,继承目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	//物品不存在
	if useIt == nil || useIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"useIndex": useIndex,
			}).Warn("goldequip:继承失败,继承消耗目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	targetItemTemplate := item.GetItemService().GetItem(int(targetIt.ItemId))
	if !targetItemTemplate.IsGoldEquip() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
				"itemId":      targetIt.ItemId,
			}).Warn("goldequip:继承失败,目标不是元神金装")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	useItemTemplate := item.GetItemService().GetItem(int(useIt.ItemId))
	if !useItemTemplate.IsGoldEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"useIndex": useIndex,
				"itemId":   useIt.ItemId,
			}).Warn("goldequip:继承失败,目标不是元神金装")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	targetPropertyData := targetIt.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	usePropertyData := useIt.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	targetUpstar := targetPropertyData.UpstarLevel
	useUpstar := usePropertyData.UpstarLevel
	if useUpstar <= targetUpstar {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
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
				"targetIndex":   targetIndex,
				"useIndex":      useIndex,
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

	//金装继承
	flag = inventoryManager.GoldEquipExtend(targetIndex, useIndex)
	if !flag {
		panic(fmt.Errorf("goldequip: goldequipExtend should be ok"))
	}

	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGoldEquipExtendBag(targetIndex, useIndex, useUpstar)
	pl.SendMsg(scMsg)
	return
}
