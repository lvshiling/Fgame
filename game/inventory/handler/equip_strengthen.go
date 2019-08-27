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
	"fgame/fgame/game/inventory/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_EQUIP_STRENGTHEN_TYPE), dispatch.HandlerFunc(handleInventoryEquipStrengthen))
}

//处理强化
func handleInventoryEquipStrengthen(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理槽位强化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryEquipStrengthen := msg.(*uipb.CSInventoryEquipStrengthen)
	slotId := csInventoryEquipStrengthen.GetSlotId()
	slotPosition := inventorytypes.BodyPositionType(slotId)
	auto := csInventoryEquipStrengthen.GetAuto()

	//参数不对
	if !slotPosition.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	typ := csInventoryEquipStrengthen.GetTyp()
	strengthenType := inventorytypes.EquipmentStrengthenType(typ)
	//参数不对
	if !strengthenType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = equipStrengthen(tpl, slotPosition, strengthenType, auto)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理槽位强化装备,错误")

		return err
	}
	log.Debug("inventory:处理槽位强化装备,完成")
	return nil
}

//强化
func equipStrengthen(pl player.Player, pos inventorytypes.BodyPositionType, strengthenType inventorytypes.EquipmentStrengthenType, auto bool) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipBag := manager.GetEquipmentBag()
	item := equipBag.GetByPosition(pos)
	//物品不存在
	if item == nil || item.IsEmpty() {
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
		return
	}

	var flag bool
	var result inventorytypes.EquipmentStrengthenResultType
	switch strengthenType {
	case inventorytypes.EquipmentStrengthenTypeStar:
		result, flag = equipmentSlotStrengthenStar(pl, pos, auto)
	case inventorytypes.EquipmentStrengthenTypeUpgrade:
		result, flag = equipmentSlotStrengthenUpgrade(pl, pos, false)
	}
	//升级不成功
	if !flag {
		return
	}
	//同步改变
	logic.SnapInventoryEquipChanged(pl)
	logic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//强化成功
	scInventoryEquipmentSlotStrength := pbutil.BuildSCInventoryEquipmentSlotStrength(pos, strengthenType, result, auto)
	pl.SendMsg(scInventoryEquipmentSlotStrength)

	return
}

//升星
func equipmentSlotStrengthenStar(pl player.Player, pos inventorytypes.BodyPositionType, auto bool) (result inventorytypes.EquipmentStrengthenResultType, flag bool) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	equipBag := manager.GetEquipmentBag()
	item := equipBag.GetByPosition(pos)
	//物品不存在
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升星失败,装备不存在")
		return
	}

	equipmentBag := manager.GetEquipmentBag()
	//判断槽位是否可以升星
	nextEquipmentStrengthenTemplate := equipmentBag.GetNextStarEquipStrengthenTemplate(pos)
	if nextEquipmentStrengthenTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升星失败,已经是满级")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotStarMax)
		return
	}

	//判断消耗条件
	items := nextEquipmentStrengthenTemplate.GetNeedItemMap()

	var hasItems map[int32]int32
	var needBuyItems map[int32]int32
	//var itemCost map[shoptypes.ShopConsumeType]int32
	silver := int64(0)
	bindGold := int32(0)
	gold := int32(0)
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	if len(items) != 0 {
		hasItems, needBuyItems = manager.GetItemsAndNeedBuy(items)
		//判断金钱
		// if len(needBuyItems) != 0 {
		// 	tItemCost, tflag := shop.GetShopService().GetShopCost(needBuyItems)
		// 	if !tflag {
		// 		log.WithFields(
		// 			log.Fields{
		// 				"playerId": pl.GetId(),
		// 				"pos":      pos.String(),
		// 			}).Warn("inventory:强化升星失败,物品不足")
		// 		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		// 		return
		// 	}
		// 	itemCost = tItemCost
		// }
		// bindGold, _ = itemCost[shoptypes.ShopConsumeTypeBindGold]
		// gold, _ = itemCost[shoptypes.ShopConsumeTypeGold]
		// silver = int64(itemCost[shoptypes.ShopConsumeTypeSliver])

		if len(needBuyItems) != 0 {
			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayerMap(pl, needBuyItems)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"pos":      pos.String(),
				}).Warn("inventory:购买物品失败,自动升星已停止")
				playerlogic.SendSystemMessage(pl, lang.ShopUpstarAutoBuyItemFail)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			gold += int32(shopNeedGold)
			bindGold += int32(shopNeedBindGold)
			silver += shopNeedSilver
		}

		tflag := propertyManager.HasEnoughCost(int64(bindGold), int64(gold), silver)
		if !tflag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"pos":      pos.String(),
				"auto":     auto,
			}).Warn("inventory:强化升星失败")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//判断是否有足够的银两
	if nextEquipmentStrengthenTemplate.SilverNum > 0 {
		if !propertyManager.HasEnoughSilver(int64(nextEquipmentStrengthenTemplate.SilverNum) + silver) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"pos":      pos.String(),
				}).Warn("inventory:强化升星失败,银两不够")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//扣除物品
	if len(hasItems) != 0 {
		reasonText := commonlog.InventoryLogReasonEquipSlotStrengthUpgrade.String()
		tflag := manager.BatchRemove(hasItems, commonlog.InventoryLogReasonEquipSlotStrengthUpgrade, reasonText)
		if !tflag {
			panic(fmt.Errorf("inventory:装备槽强化升星移除材料应该成功"))
		}
	}

	if len(needBuyItems) != 0 {
		goldReason := commonlog.GoldLogReasonEquipSlotStrengthUpgradeAuto
		goldReasonText := fmt.Sprintf(goldReason.String(), pos.String(), nextEquipmentStrengthenTemplate.Level)
		silverReason := commonlog.SilverLogReasonEquipSlotStrengthUpgradeAuto
		silverReasonText := fmt.Sprintf(silverReason.String(), pos.String(), nextEquipmentStrengthenTemplate.Level)
		tflag := propertyManager.Cost(int64(bindGold), int64(gold), goldReason, goldReasonText, silver, silverReason, silverReasonText)
		if !tflag {
			panic(fmt.Errorf("inventory:装备槽强化升星自动购买应该成功"))
		}
	}

	if nextEquipmentStrengthenTemplate.SilverNum > 0 {
		reasonText := commonlog.SilverLogReasonEquipSlotStrengthUpgrade.String()
		tflag := propertyManager.CostSilver(int64(nextEquipmentStrengthenTemplate.SilverNum), commonlog.SilverLogReasonEquipSlotStrengthUpgrade, reasonText)
		if !tflag {
			panic(fmt.Errorf("inventory:装备槽强化升星花费银两应该成功"))
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//判断是否可以升星
	success := mathutils.RandomHit(common.MAX_RATE, int(nextEquipmentStrengthenTemplate.SuccessRate))
	if success {
		//升星
		tflag := equipmentBag.StrengthStar(pos)
		if !tflag {
			panic(fmt.Errorf("inventory: strength star should be ok"))
		}
		flag = true
		result = inventorytypes.EquipmentStrengthenResultTypeSuccess
		inventorylogic.UpdateEquipmentProperty(pl)

		return
	}
	//判断是否会回退
	if nextEquipmentStrengthenTemplate.GetFailEquipStrengthenTemplate() != nil {
		fail := mathutils.RandomHit(common.MAX_RATE, int(nextEquipmentStrengthenTemplate.ReturnRate))
		if fail {
			//回退
			tflag := equipmentBag.StrengthStarBack(pos)
			if !tflag {
				panic(fmt.Errorf("inventory: strength star back should be ok"))
			}
			flag = true
			result = inventorytypes.EquipmentStrengthenResultTypeBack
			return
		}
	}

	flag = true
	result = inventorytypes.EquipmentStrengthenResultTypeFailed
	return
}

//强化升级
func equipmentSlotStrengthenUpgrade(pl player.Player, pos inventorytypes.BodyPositionType, ignoreErrorMsg bool) (result inventorytypes.EquipmentStrengthenResultType, flag bool) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	item := manager.GetEquipByPos(pos)
	//物品不存在
	if item == nil || item.IsEmpty() {
		if ignoreErrorMsg {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升级失败,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	equipmentBag := manager.GetEquipmentBag()

	//判断是否达到上限
	// if !equipmentBag.IfCanStrengthLevel(pos) {
	// 	if ignoreErrorMsg {
	// 		return
	// 	}
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"pos":      pos.String(),
	// 		}).Warn("inventory:强化升级失败,已经满级")
	// 	playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotStarMax)
	// 	return
	// }
	//判断槽位是否可以升星
	nextEquipmentStrengthenTemplate := equipmentBag.GetNextUpgradeEquipStrengthenTemplate(pos)
	if nextEquipmentStrengthenTemplate == nil {
		if ignoreErrorMsg {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotStarMax)
		return
	}
	maxLevel := int32(math.Floor(float64(pl.GetLevel()) / float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEquipmentStrengthenLevelLimit))))
	//判断等级限制
	if nextEquipmentStrengthenTemplate.Level > maxLevel {
		if ignoreErrorMsg {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升级失败,达到极限")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotLevelExceedLevel)
		return
	}

	items := nextEquipmentStrengthenTemplate.GetNeedItemMap()
	if len(items) != 0 {
		if !manager.HasEnoughItems(items) {
			if ignoreErrorMsg {
				return
			}
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"pos":      pos.String(),
				}).Warn("inventory:强化升级失败,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}
	//判断是否有足够的银两
	if nextEquipmentStrengthenTemplate.SilverNum > 0 {
		if !propertyManager.HasEnoughSilver(int64(nextEquipmentStrengthenTemplate.SilverNum)) {
			if ignoreErrorMsg {
				return
			}
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"pos":      pos.String(),
				}).Warn("inventory:强化升级失败,银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	if len(items) != 0 {
		reasonText := commonlog.InventoryLogReasonEquipSlotStrengthUpgrade.String()
		flag := manager.BatchRemove(items, commonlog.InventoryLogReasonEquipSlotStrengthUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("inventory:装备槽强化升级移除材料应该成功"))
		}
	}
	if nextEquipmentStrengthenTemplate.SilverNum > 0 {
		reasonText := commonlog.SilverLogReasonEquipSlotStrengthUpgrade.String()
		flag := propertyManager.CostSilver(int64(nextEquipmentStrengthenTemplate.SilverNum), commonlog.SilverLogReasonEquipSlotStrengthUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("inventory:装备槽强化升级花费银两应该成功"))
		}
	}

	success := mathutils.RandomHit(common.MAX_RATE, int(nextEquipmentStrengthenTemplate.SuccessRate))
	if success {
		//升星
		flag = equipmentBag.StrengthLevel(pos)
		if !flag {
			panic(fmt.Errorf("inventory: 强化升级应该成功"))
		}

		inventorylogic.UpdateEquipmentProperty(pl)

		result = inventorytypes.EquipmentStrengthenResultTypeSuccess
		return
	}

	//判断是否会回退
	if nextEquipmentStrengthenTemplate.GetFailEquipStrengthenTemplate() != nil {
		fail := mathutils.RandomHit(common.MAX_RATE, int(nextEquipmentStrengthenTemplate.ReturnRate))
		if fail {
			//回退
			flag = equipmentBag.StrengthLevelBack(pos)
			if !flag {
				panic(fmt.Errorf("inventory: strength star back should be ok"))
			}
			result = inventorytypes.EquipmentStrengthenResultTypeBack
			return
		}
	}

	flag = true
	result = inventorytypes.EquipmentStrengthenResultTypeFailed
	return
}
