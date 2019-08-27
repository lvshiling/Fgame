package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_UPSTAR_BODY_TYPE), dispatch.HandlerFunc(handleGoldEquipUpstarBody))
}

//处理升星强化
func handleGoldEquipUpstarBody(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理装备槽金装升星强化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGoldEquipUpstarBody)
	targetSlotId := csMsg.GetSlotId()
	isAvoid := csMsg.GetIsAvoid()

	posType := inventorytypes.BodyPositionType(targetSlotId)
	if !posType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetSlotId,
				"error":       err,
			}).Warn("goldequip:处理装备槽金装升星强化,位置错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = goldEquipUpstarBoday(tpl, posType, isAvoid)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,

				"error": err,
			}).Error("goldequip:处理装备槽金装升星强化,错误")

		return err
	}
	log.Debug("goldequip:处理装备槽金装升星强化,完成")
	return nil
}

//升星强化
func goldEquipUpstarBoday(pl player.Player, posType inventorytypes.BodyPositionType, isAvoid bool) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	equipBag := goldequipManager.GetGoldEquipBag()
	targetIt := equipBag.GetByPosition(posType)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,
			}).Warn("goldequip:升星强化升级失败,升星强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	targetItemTemplate := item.GetItemService().GetItem(int(targetIt.GetItemId()))
	propertyData := targetIt.GetPropertyData().(*goldequiptypes.GoldEquipPropertyData)

	//品质
	if targetItemTemplate.GetQualityType() < itemtypes.ItemQualityTypeBlue {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,
			}).Warn("goldequip:升星强化升级失败,仅橙色品质元神装备可以进行强化升星")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipUpstarNotAllow)
		return
	}

	//能否被升星强化
	goldEquipTemplate := targetItemTemplate.GetGoldEquipTemplate()
	if goldEquipTemplate.GoldeuipUpstarId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,
			}).Warn("goldequip:升星强化升级失败,该金装无法被升星强化")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipUpstarNotAllow)
		return
	}

	//升星强化等级判断
	upstarTemp := goldEquipTemplate.GetUpstarTemplate(propertyData.UpstarLevel)
	if upstarTemp.NextId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,
			}).Warn("goldequip:升星强化升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipReachUpstarFullLevel)
		return
	}

	//
	nextUpstarTemp := upstarTemp.GetNextTemplate()
	if isAvoid && nextUpstarTemp.ProtectItemId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,
			}).Warn("goldequip:升星强化升级失败,该物品不能使用防爆符")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	needItemMap := nextUpstarTemp.GetNeedItemMap()
	if isAvoid {
		needItemMap[nextUpstarTemp.ProtectItemId] = nextUpstarTemp.ProtectItemCount
	}
	if !inventoryManager.HasEnoughItems(needItemMap) {
		log.WithFields(
			log.Fields{
				"playerid":    pl.GetId(),
				"posType":     posType,
				"needItemMap": needItemMap,
			}).Warn("goldequip:升星强化升级失败,道具不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//消耗材料
	if len(needItemMap) != 0 {
		useReason := commonlog.InventoryLogReasonGoldEquipUpstar
		flag := inventoryManager.BatchRemove(needItemMap, useReason, useReason.String())
		if !flag {
			panic(fmt.Errorf("goldequip:装备槽升星强化升级移除材料应该成功"))
		}
	}

	var result goldequiptypes.UpstarResultType
	isReturn := false
	//计算成功
	isSuccess := mathutils.RandomHit(common.MAX_RATE, int(nextUpstarTemp.SuccessRate))
	if isSuccess {
		flag := equipBag.UpstarSuccess(posType)
		if !flag {
			panic(fmt.Errorf("goldequip: 升星强化升级应该成功"))
		}
		result = goldequiptypes.UpstarResultTypeSuccess
	} else {
		isBack := mathutils.RandomHit(common.MAX_RATE, int(nextUpstarTemp.FailReturnRate))
		if isBack && !isAvoid {
			isReturn = true
			returnLevel := nextUpstarTemp.GetFaildReturnTemplate().Level
			equipBag.UpstarFaildReturn(posType, returnLevel)
			result = goldequiptypes.UpstarResultTypeBack
		} else {
			result = goldequiptypes.UpstarResultTypeFailed
		}
	}

	if isSuccess || isReturn {
		goldequiplogic.GoldEquipPropertyChanged(pl)
		propertylogic.SnapChangedProperty(pl)
	}

	//同步改变
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGoldEquipUpstarBody(result, propertyData.UpstarLevel, posType, isAvoid)
	pl.SendMsg(scMsg)
	return
}
