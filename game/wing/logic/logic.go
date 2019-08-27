package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
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
	wingeventtypes "fgame/fgame/game/wing/event/types"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"
	"fgame/fgame/game/wing/wing"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//战翼进阶的逻辑
func HandleWingAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeWing)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("wing:战翼升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	beforeNum := int32(wingInfo.AdvanceId)
	nextAdvancedId := wingInfo.AdvanceId + 1
	wingTemplate := wing.GetWingService().GetWingNumber(int32(nextAdvancedId))
	if wingTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("Wing:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.WingAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := wingTemplate.UseMoney
	//进阶需要消耗的银两
	costSilver := int64(wingTemplate.UseYinliang)
	//进阶需要消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := wingTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := wingTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = wingTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("Wing:战翼进阶丹不足,无法进阶")
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
		// costBindGold += shopNeedBindGold
		// costGold += shopNeedGold
		// costSilver += shopNeedSilver

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("wing:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("wing:购买物品失败,自动进阶已停止")
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
			}).Warn("wing:银两不足,无法进阶")
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
			}).Warn("wing:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够元宝
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("wing:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonWingAdvanced.String()
	reasonSliverText := commonlog.SilverLogReasonWingAdvanced.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonWingAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonWingAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("wing: wingAdvanced Cost should be ok"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonWingAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), wingInfo.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("wing: wingAdvanced use item shoud be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	if useItemTemplate != nil {
		//消耗事件
		gameevent.Emit(wingeventtypes.EventTypeWingAdvancedCost, pl, int32(wingInfo.AdvanceId))
	}

	//战翼进阶判断
	sucess, pro, randBless, addTimes, isDouble := WingAdvanced(pl, wingInfo.TimesNum, wingInfo.Bless, wingTemplate)
	wingManager.WingAdvanced(pro, addTimes, sucess)
	if sucess {
		wingReason := commonlog.WingLogReasonAdvanced
		reasonText := fmt.Sprintf(wingReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := wingeventtypes.CreatePlayerWingAdvancedLogEventData(beforeNum, 1, wingReason, reasonText)
		gameevent.Emit(wingeventtypes.EventTypeWingAdvancedLog, pl, data)

		//同步属性
		WingPropertyChanged(pl)
		advancedFinish(pl, int32(wingInfo.AdvanceId), wingInfo.WingId)
	} else {
		//进阶不成功
		advancedBless(pl, wingInfo.AdvanceId, wingInfo.WingId, randBless, wingInfo.Bless, wingInfo.BlessTime, isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32, wingId int32) (err error) {
	scWingAdvanced := pbutil.BuildSCWingAdavancedFinshed(int32(advancedId), wingId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scWingAdvanced)
	return
}

func advancedBless(pl player.Player, advancedId int, wingId int32, bless int32, totalBless int32, blessTime int64, isDouble bool) (err error) {
	scWingAdvanced := pbutil.BuildSCWingAdavanced(int32(advancedId), wingId, bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scWingAdvanced)
	return
}

//变更战翼属性
func WingPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeWing.Mask())
	return
}

//护体仙羽进阶的逻辑
func HandleFeatherAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeFeather)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("wing:护体仙羽升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	WingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	WingInfo := WingManager.GetWingInfo()
	beforeNum := WingInfo.FeatherId
	nextAdvancedId := WingInfo.FeatherId + 1
	featherTemplate := wing.GetWingService().GetFeather(nextAdvancedId)
	if featherTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("Wing:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.WingAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := featherTemplate.UseGold
	//进阶需要消耗绑元
	costBindGold := featherTemplate.UseBindGold
	//进阶需要消耗的银两
	costSilver := int64(featherTemplate.UseSilver)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := featherTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := featherTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = featherTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("Wing:护体仙羽进阶丹不足,无法进阶")
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
		// 	}).Warn("wing:商铺没有该道具,无法自动购买")
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
				}).Warn("wing:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("wing:购买物品失败,自动进阶已停止")
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
			}).Warn("wing:银两不足,无法进阶")
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
			}).Warn("wing:元宝不足,无法进阶")
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
			}).Warn("wing:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonFeatherAdvanced.String()
	reasonSliverText := commonlog.SilverLogReasonFeatherAdvanced.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonFeatherAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonFeatherAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("wing: featherAdvanced Cost should be ok"))
	}
	//同步元宝
	if costBindGold != 0 || costGold != 0 || costSilver != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonFeatherAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), WingInfo.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("wing: featherAdvanced use item shoud be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//护体仙羽进阶判断
	pro, _, sucess := FeatherAdvanced(pl, WingInfo.FeatherNum, WingInfo.FeatherPro, featherTemplate)
	WingManager.FeatherFeed(pro, 1, sucess)
	if sucess {
		featherReason := commonlog.WingLogReasonFeatherAdvanced
		reasonText := fmt.Sprintf(featherReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := wingeventtypes.CreatePlayerFeatherAdvancedLogEventData(beforeNum, 1, featherReason, reasonText)
		gameevent.Emit(wingeventtypes.EventTypeFeatherAdvancedLog, pl, data)

		//同步属性
		FeatherPropertyChanged(pl)
	}
	scFeatherAdvanced := pbutil.BuildSCFeatherAdvanced(WingInfo.FeatherId, WingInfo.FeatherPro, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scFeatherAdvanced)
	return
}

//变更护体仙羽属性
func FeatherPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeFeather.Mask())
	return
}

//护体仙羽进阶判断
func FeatherAdvanced(pl player.Player, curTimesNum int32, curBless int32, featherTemplate *gametemplate.FeatherTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeFeather, featherTemplate.TimesMin, featherTemplate.TimesMax)
	updateRate := featherTemplate.UpdateWfb
	blessMax := featherTemplate.NeedRate
	addMin := featherTemplate.AddMin
	addMax := featherTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//战翼进阶判断
func WingAdvanced(pl player.Player, curTimesNum int32, curBless int32, wingTemplate *gametemplate.WingTemplate) (sucess bool, pro int32, randBless int32, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeWing, wingTemplate.TimesMin, wingTemplate.TimesMax)
	updateRate := wingTemplate.UpdateWfb
	blessMax := wingTemplate.ZhufuMax
	addMin := wingTemplate.AddMin
	addMax := wingTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeWing)
	if isDouble {
		addTimes += extralAddTimes
		randBless += wingTemplate.AddMax
	}
	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//战翼祝福丹
func WingEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, wingTemplate *gametemplate.WingTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeWing, wingTemplate.TimesMin, wingTemplate.TimesMax)
	updateRate := wingTemplate.UpdateWfb
	blessMax := wingTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//仙羽祝福丹
func FeatherEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, featherTemplate *gametemplate.FeatherTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeFeather, featherTemplate.TimesMin, featherTemplate.TimesMax)
	updateRate := featherTemplate.UpdateWfb
	blessMax := featherTemplate.NeedRate

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//战翼幻化培养判断
func WingHuanHua(curTimesNum int32, curBless int32, huanHuaTemplate *gametemplate.WingHuanHuaTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := huanHuaTemplate.TimesMin
	timesMax := huanHuaTemplate.TimesMax
	updateRate := huanHuaTemplate.UpdateWfb
	blessMax := huanHuaTemplate.ZhufuMax
	addMin := huanHuaTemplate.AddMin
	addMax := huanHuaTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//战翼皮肤升星判断
func WingSkinUpstar(curTimesNum int32, curBless int32, upstarTemplate *gametemplate.WingUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upstarTemplate.TimesMin
	timesMax := upstarTemplate.TimesMax
	updateRate := upstarTemplate.UpstarRate
	blessMax := upstarTemplate.ZhufuMax
	addMin := upstarTemplate.AddMin
	addMax := upstarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
