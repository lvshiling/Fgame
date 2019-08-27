package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/goldequip/pbutil"
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
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_UPSTAR_BAG_TYPE), dispatch.HandlerFunc(handleGoldEquipUpstarBag))
}

//处理背包金装升星强化
func handleGoldEquipUpstarBag(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理背包金装升星强化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGoldEquipUpstarBag)
	targetIndex := csMsg.GetItemIdex()
	isAvoid := csMsg.GetIsAvoid()

	err = goldEquipUpstarBag(tpl, targetIndex, isAvoid)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
				"error":       err,
			}).Error("goldequip:处理背包金装升星强化,错误")

		return err
	}
	log.Debug("goldequip:处理背包金装升星强化,完成")
	return nil
}

//升星强化
func goldEquipUpstarBag(pl player.Player, targetIndex int32, isAvoid bool) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	targetIt := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, targetIndex)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:升星强化升级失败,升星强化目标不存在")
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
			}).Warn("goldequip:强化升级失败,目标不是元神金装")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	propertyData := targetIt.PropertyData.(*goldequiptypes.GoldEquipPropertyData)

	//品质
	if targetItemTemplate.GetQualityType() < itemtypes.ItemQualityTypeBlue {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:升星强化升级失败,仅蓝色以及之上品质元神装备可以进行强化")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipUpstarNotAllow)
		return
	}

	//能否被升星强化
	goldEquipTemplate := targetItemTemplate.GetGoldEquipTemplate()
	if goldEquipTemplate.GoldeuipUpstarId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:升星强化升级失败,该金装无法被升星强化")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipUpstarNotAllow)
		return
	}

	//升星强化等级判断
	upstarTemp := goldEquipTemplate.GetUpstarTemplate(propertyData.UpstarLevel)
	if upstarTemp.NextId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIt.Index,
				"itemId":      targetIt.ItemId,
				"curLevel":    propertyData.UpstarLevel,
			}).Warn("goldequip:升星强化升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipReachUpstarFullLevel)
		return
	}

	//
	nextUpstarTemp := upstarTemp.GetNextTemplate()
	if isAvoid && nextUpstarTemp.ProtectItemId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
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
				"targetIndex": targetIndex,
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
		flag := inventoryManager.GoldEquipUpstarSuccess(targetIndex)
		if !flag {
			panic(fmt.Errorf("goldequip: 升星强化升级应该成功"))
		}
		result = goldequiptypes.UpstarResultTypeSuccess
	} else {
		// 回退计算
		isReturn = mathutils.RandomHit(common.MAX_RATE, int(nextUpstarTemp.FailReturnRate))
		if isReturn && !isAvoid {
			returnLevel := nextUpstarTemp.GetFaildReturnTemplate().Level
			inventoryManager.GoldEquipUpstarReturn(targetIndex, returnLevel)
			result = goldequiptypes.UpstarResultTypeBack
		} else {
			result = goldequiptypes.UpstarResultTypeFailed
		}
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGoldEquipUpstarBag(result, propertyData.UpstarLevel, targetIndex, isAvoid)
	pl.SendMsg(scMsg)
	return
}
