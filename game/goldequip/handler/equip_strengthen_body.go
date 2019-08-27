package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_STRENGTHEN_BODY_TYPE), dispatch.HandlerFunc(handleGoldEquipStrengthenBody))
}

//处理强化
func handleGoldEquipStrengthenBody(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理装备槽金装强化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGoldEquipStrengthenBody)
	targetSlotId := csMsg.GetTargetSlotId()
	useItemNum := csMsg.GetItemNum()

	if useItemNum < 1 {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
				"UseItemNum":   useItemNum,
			}).Warn("goldequip:强化升级失败,材料数量不能小于1")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	posType := inventorytypes.BodyPositionType(targetSlotId)
	if !posType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetSlotId,
				"useItemNum":  useItemNum,
			}).Warn("goldequip:处理装备槽金装强化,位置错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = goldEquipStrengthenBoday(tpl, posType, useItemNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
				"useItemNum":   useItemNum,
				"error":        err,
			}).Error("goldequip:处理装备槽金装强化,错误")

		return err
	}
	log.Debug("goldequip:处理装备槽金装强化,完成")
	return nil
}

//强化
func goldEquipStrengthenBoday(pl player.Player, posType inventorytypes.BodyPositionType, useItemNum int32) (err error) {

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	equipBag := goldequipManager.GetGoldEquipBag()
	targetIt := equipBag.GetByPosition(posType)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"posType":    posType,
				"useItemNum": useItemNum,
			}).Warn("goldequip:强化升级失败,强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	curLevel := targetIt.GetLevel()
	targetItemTemplate := item.GetItemService().GetItem(int(targetIt.GetItemId()))

	//品质
	if targetItemTemplate.GetQualityType() < itemtypes.ItemQualityTypeOrange {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType,
			}).Warn("goldequip:强化升级失败,仅橙色品质元神装备可以进行强化")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipStrengthenQualityNotEnough)
		return
	}

	//能否被强化
	goldEquipTemplate := targetItemTemplate.GetGoldEquipTemplate()
	if goldEquipTemplate.GoldequipStrenId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"posType":    posType,
				"useItemNum": useItemNum,
			}).Warn("goldequip:强化升级失败,该金装无法被强化")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipStrengthenNotAllow)
		return
	}

	//强化等级判断
	strengthenTemp := goldEquipTemplate.GetStrengthenTemplate(curLevel)
	if strengthenTemp.NextId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"posType":    posType,
				"useItemNum": useItemNum,
			}).Warn("goldequip:强化升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipReachStrengthenMax)
		return
	}

	//材料数量
	maxNeedItemNum := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGoldEquipStrengthenItemUseMax))
	if useItemNum > maxNeedItemNum {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"maxNeedItemNum": maxNeedItemNum,
				"useItemNum":     useItemNum,
			}).Warn("inventory:强化失败,材料数量超出限制")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipGreathenStrengthenItemUseMax)
		return
	}

	needItemMap := strengthenTemp.GetNeedItemMap(useItemNum)
	if !inventoryManager.HasEnoughItems(needItemMap) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"UseItemNum": useItemNum,
			}).Warn("goldequip:强化升级失败,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//成功率
	totalRate := goldequiplogic.CountGoldEquipStrengthenRate(curLevel, needItemMap)

	//消耗材料
	if len(needItemMap) > 0 {
		reason := commonlog.InventoryLogReasonGoldEquipStrengthUpgrade
		flag := inventoryManager.BatchRemove(needItemMap, reason, reason.String())
		if !flag {
			panic(fmt.Errorf("goldequip:背包强化升级移除材料应该成功"))
		}
	}

	//计算成功
	success := mathutils.RandomHit(common.MAX_RATE, int(totalRate))
	if success {
		flag := equipBag.UpdateGoldEquipLevel(posType)
		if !flag {
			panic(fmt.Errorf("goldequip: 强化升级应该成功"))
		}

		goldequiplogic.GoldEquipPropertyChanged(pl)
		// propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeGoldequipPercent.Mask())
	}

	//同步改变
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scGoldEquipStrengthen := pbutil.BuildSCGoldEquipStrengthenBody(posType, success)
	pl.SendMsg(scGoldEquipStrengthen)

	return
}
