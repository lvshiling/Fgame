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
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_STRENGTHEN_BUWEI), dispatch.HandlerFunc(handleGoldEquipStrengthenBuWei))
}

//处理新强化部位
func handleGoldEquipStrengthenBuWei(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理装备槽金装新强化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGoldEquipStrengthenBuwei)
	targetSlotId := csMsg.GetSlotId()
	isAvoid := csMsg.GetIsAvoid()

	posType := inventorytypes.BodyPositionType(targetSlotId)
	if !posType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetSlotId,
				"error":       err,
			}).Warn("goldequip:处理装备槽金装新强化,位置错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = goldEquipStrengthenBuWei(tpl, posType, isAvoid)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,

				"error": err,
			}).Error("goldequip:处理装备槽金装新强化,错误")

		return err
	}
	log.Debug("goldequip:处理装备槽金装新强化,完成")
	return nil
}

//新强化
func goldEquipStrengthenBuWei(pl player.Player, posType inventorytypes.BodyPositionType, isAvoid bool) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	equipBag := goldequipManager.GetGoldEquipBag()
	targetIt := equipBag.GetByPosition(posType)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType.String(),
			}).Warn("goldequip:新强化升级失败,新强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	nextTemp := equipBag.GetNextStrengthenBuWeiTemplate(posType)
	if nextTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      posType.String(),
			}).Warn("goldequip:新强化升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotLevelMax)
		return
	}

	if isAvoid && nextTemp.ProtectItemId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,
			}).Warn("goldequip:新强化升级失败,该物品不能使用防爆符")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	needItemMap := nextTemp.GetNeedItemMap()
	if isAvoid {
		needItemMap[nextTemp.ProtectItemId] = nextTemp.ProtectItemCount
	}
	if !inventoryManager.HasEnoughItems(needItemMap) {
		log.WithFields(
			log.Fields{
				"playerid":    pl.GetId(),
				"posType":     posType.String(),
				"needItemMap": needItemMap,
			}).Warn("goldequip:新强化升级失败,道具不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//消耗材料
	if len(needItemMap) != 0 {
		useReason := commonlog.InventoryLogReasonGoldEquipUpstar
		flag := inventoryManager.BatchRemove(needItemMap, useReason, useReason.String())
		if !flag {
			panic(fmt.Errorf("goldequip:装备槽新强化升级移除材料应该成功"))
		}
	}

	var result goldequiptypes.UpstarResultType
	isReturn := false
	//计算成功
	isSuccess := mathutils.RandomHit(common.MAX_RATE, int(nextTemp.SuccessRate))
	if isSuccess {
		flag := equipBag.StrengthBuWeiLevelUp(posType)
		if !flag {
			panic(fmt.Errorf("goldequip: 新强化升级应该成功"))
		}
		result = goldequiptypes.UpstarResultTypeSuccess
	} else {
		isBack := mathutils.RandomHit(common.MAX_RATE, int(nextTemp.FailReturnRate))
		if isBack && !isAvoid {
			isReturn = true
			returnLevel := nextTemp.GetFaildReturnTemplate().Level
			equipBag.StrengthBuWeiLevelReturn(posType, returnLevel)
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

	scMsg := pbutil.BuildSCGoldEquipStrengthenBuwei(result, targetIt.GetNewStLevel(), posType, isAvoid)
	pl.SendMsg(scMsg)
	return
}
