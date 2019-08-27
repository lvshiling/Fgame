package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/bodyshield/bodyshield"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	"fgame/fgame/game/bodyshield/pbutil"
	playerbshield "fgame/fgame/game/bodyshield/player"
	commomlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//护体盾进阶
func HandleBodyShieldAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeBodyShield)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("bodyshield:护体盾升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	bshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
	bodyShieldInfo := bshieldManager.GetBodyShiedInfo()
	nextAdvancedId := bodyShieldInfo.AdvanceId + 1
	bodyShieldTemplate := bodyshield.GetBodyShieldService().GetBodyShieldNumber(int32(nextAdvancedId))
	if bodyShieldTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Warn("bodyshield:神盾尖刺已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.BodyShieldShieldReachedLimits)
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//进阶需要消耗的元宝
	costGold := bodyShieldTemplate.UseMoney
	//进阶需要消耗的银两
	costSilver := int64(bodyShieldTemplate.UseYinliang)
	//进阶需要的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := bodyShieldTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := bodyShieldTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = bodyShieldTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("bodyshield:神盾尖刺进阶丹不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//自动进阶
		needBuyNum := itemCount - totalNum
		itemCount = totalNum
		//获取价格
		// shopTemplate := shop.GetShopService().GetShopTemplateByItem(useItem)
		// if shopTemplate == nil {
		// 	log.WithFields(log.Fields{
		// 		"playerId": pl.GetId(),
		// 	}).Warn("bodyshield:商铺没有该道具,无法自动购买")
		// 	playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
		// 	return
		// }
		// shopNeedGold, shopNeedBindGold, shopNeedSilver := shopTemplate.GetConsumeData(needBuyNum)
		// costGold += shopNeedGold
		// costBindGold += shopNeedBindGold
		// costSilver += shopNeedSilver

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("bodyshield:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("bodyshield:购买物品失败,自动进阶已停止")
				playerlogic.SendSystemMessage(pl, lang.ShopAdvancedAutoBuyItemFail)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			costGold += int32(shopNeedGold)
			costBindGold += int32(shopNeedBindGold)
			costSilver += shopNeedSilver
		}
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("bodyshield:银两不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("bodyshield:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("mount:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonBodyShieldAdvanced.String()
	reasonSliverText := commonlog.SilverLogReasonBodyShieldAdvanced.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonBodyShieldAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonBodyShieldAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("bodyshield: bodyShieldAdvanced Cost should be ok"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonShieldAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), bodyShieldInfo.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("bodyshield: bodyShieldAdvanced use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)

	}
	if useItemTemplate != nil {
		//消耗事件
		gameevent.Emit(bodyshieldeventtypes.EventTypeBodyShieldAdvancedCost, pl, int32(bodyShieldInfo.AdvanceId))
	}

	//进阶判断
	beforeNum := int32(bodyShieldInfo.AdvanceId)
	sucess, pro, randBless, addTimes, isDouble := BodyShieldAdvanced(pl, bodyShieldInfo.TimesNum, bodyShieldInfo.Bless, bodyShieldTemplate)
	bshieldManager.BodyShieldAdvanced(pro, addTimes, sucess)
	if sucess {
		//同步属性
		BodyShieldPropertyChanged(pl)
		advancedFinish(pl, int32(bodyShieldInfo.AdvanceId))

		shieldReason := commonlog.ShieldLogReasonBodyAdvanced
		reasonText := fmt.Sprintf(shieldReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := bodyshieldeventtypes.CreatePlayerBodyShieldAdvancedLogEventData(beforeNum, 1, shieldReason, reasonText)
		gameevent.Emit(bodyshieldeventtypes.EventTypeBodyShieldAdvancedLog, pl, data)
	} else {
		//进阶不成功
		advancedBless(pl, bodyShieldInfo.AdvanceId, randBless, bodyShieldInfo.Bless, bodyShieldInfo.BlessTime, isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32) {
	scBodyShieldAdvanced := pbutil.BuildSCBodyShieldAdavancedFinshed(int32(advancedId), commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scBodyShieldAdvanced)
	return
}

func advancedBless(pl player.Player, advancedId int, bless int32, totalBless int32, blessTime int64, isDouble bool) {
	scBodyShieldAdvanced := pbutil.BuildSCBodyShieldAdavanced(int32(advancedId), bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scBodyShieldAdvanced)
	return
}

//变更护体盾属性
func BodyShieldPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeBodyShield.Mask())
	return
}

//变更神盾尖刺属性
func ShieldPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeShield.Mask())
	return
}

//护体盾进阶判断
func BodyShieldAdvanced(pl player.Player, curTimesNum int32, curBless int32, bodyShieldTemplate *gametemplate.BodyShieldTemplate) (sucess bool, pro int32, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeBodyShield, bodyShieldTemplate.TimesMin, bodyShieldTemplate.TimesMax)
	updateRate := bodyShieldTemplate.UpdateWfb
	blessMax := bodyShieldTemplate.ZhufuMax
	addMin := bodyShieldTemplate.AddMin
	addMax := bodyShieldTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeBodyshield)
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += bodyShieldTemplate.AddMax
	}
	curTimesNum += addTimes

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//护体盾祝福丹
func BodyShieldEatZhuFuDan(pl player.Player, curTimesNum, curBless, randBless int32, bodyShieldTemplate *gametemplate.BodyShieldTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeBodyShield, bodyShieldTemplate.TimesMin, bodyShieldTemplate.TimesMax)
	updateRate := bodyShieldTemplate.UpdateWfb
	blessMax := bodyShieldTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//金甲丹培养判断
func BodyShieldJinJiaDan(curTimesNum int32, curBless int32, jinJiaDanTemplate *gametemplate.BodyShieldJinJiaTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := jinJiaDanTemplate.TimesMin
	timesMax := jinJiaDanTemplate.TimesMax
	updateRate := jinJiaDanTemplate.UpdateWfb
	blessMax := jinJiaDanTemplate.ZhufuMax
	addMin := jinJiaDanTemplate.AddMin
	addMax := jinJiaDanTemplate.AddMax + 1

	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//神盾尖刺进阶的逻辑
func HandleShieldAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeShield)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("bodyshield:盾刺升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
	bshield := manager.GetBodyShiedInfo()
	nextAdvancedId := bshield.ShieldId + 1
	shieldTemplate := bodyshield.GetBodyShieldService().GetShield(nextAdvancedId)
	if shieldTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Warn("bodyshield:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.BodyShieldAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := shieldTemplate.UseGold
	//进阶需要消耗绑元
	costBindGold := shieldTemplate.UseBindGold
	//进阶需要消耗的银两
	costSilver := int64(shieldTemplate.UseSilver)

	//需要消耗物品
	itemCount := int32(0)
	useItem := shieldTemplate.UseItem
	totalNum := int32(0)
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := shieldTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = shieldTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("bodyshield:神盾尖刺进阶丹不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动进阶
		needBuyNum := itemCount - totalNum
		itemCount = totalNum
		//获取价格
		// shopTemplate := shop.GetShopService().GetShopTemplateByItem(useItem)
		// if shopTemplate == nil {
		// 	log.WithFields(log.Fields{
		// 		"playerId": pl.GetId(),
		// 		"autoFlag": autoFlag,
		// 	}).Warn("bodyshield:商铺没有该道具,无法自动购买")
		// 	playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
		// 	return
		// }
		// shopNeedGold, shopNeedBindGold, shopNeedSilver := shopTemplate.GetConsumeData(needBuyNum)
		// costGold += shopNeedGold
		// costBindGold += shopNeedBindGold
		// costSilver += shopNeedSilver

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("bodyshield:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("bodyshield:购买物品失败,自动进阶已停止")
				playerlogic.SendSystemMessage(pl, lang.ShopAdvancedAutoBuyItemFail)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			costGold += int32(shopNeedGold)
			costBindGold += int32(shopNeedBindGold)
			costSilver += shopNeedSilver
		}
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("bodyshield:银两不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("bodyshield:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("bodyshield:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonShieldAdvanced.String()
	reasonSliverText := commonlog.SilverLogReasonShieldAdvanced.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonShieldAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonShieldAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("bodyshield: shieldAdvanced Cost should be ok"))
	}
	//同步元宝
	if costBindGold != 0 || costGold != 0 || costSilver != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonShieldSpikesAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), bshield.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("bodyshield: shieldAdvanced use item shoud be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)

	}
	if useItemTemplate != nil {
		//消耗事件
		gameevent.Emit(bodyshieldeventtypes.EventTypeShieldAdvancedCost, pl, bshield.ShieldId)
	}

	//培养判断
	beforeNum := int32(bshield.AdvanceId)
	sucess, pro, addTimes, isDouble := ShieldAdvanced(pl, bshield.ShieldNum, bshield.ShieldPro, shieldTemplate)
	manager.ShieldFeed(pro, addTimes, sucess)
	//同步属性
	if sucess {
		shieldReason := commonlog.ShieldLogReasonAdvanced
		reasonText := fmt.Sprintf(shieldReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := bodyshieldeventtypes.CreatePlayerShieldAdvancedLogEventData(beforeNum, 1, shieldReason, reasonText)
		gameevent.Emit(bodyshieldeventtypes.EventTypeShieldAdvancedLog, pl, data)

		ShieldPropertyChanged(pl)
	}
	scShieldAdvanced := pbutil.BuildSCShieldAdvanced(bshield.ShieldId, bshield.ShieldPro, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scShieldAdvanced)
	return
}

//神盾尖刺进阶判断
func ShieldAdvanced(pl player.Player, curTimesNum int32, curBless int32, shieldTemplate *gametemplate.ShieldTemplate) (sucess bool, pro int32, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeShield, shieldTemplate.TimesMin, shieldTemplate.TimesMax)
	updateRate := shieldTemplate.UpdateWfb
	blessMax := shieldTemplate.NeedRate
	addMin := shieldTemplate.AddMin
	addMax := shieldTemplate.AddMax + 1

	randBless := int32(mathutils.RandomRange(int(addMin), int(addMax)))
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeShield)
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += shieldTemplate.AddMax
	}
	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//神盾尖刺祝福丹判断
func ShieldEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, shieldTemplate *gametemplate.ShieldTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeShield, shieldTemplate.TimesMin, shieldTemplate.TimesMax)
	updateRate := shieldTemplate.UpdateWfb
	blessMax := shieldTemplate.NeedRate

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}
