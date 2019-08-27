package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
	commomlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
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

//变更暗器属性
func AnqiPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeAnqi.Mask())
	return
}

//暗器进阶判断
func AnqiAdvanced(pl player.Player, curTimesNum int32, curBless int32, anqiTemplate *gametemplate.AnqiTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeAnqi, anqiTemplate.TimesMin, anqiTemplate.TimesMax)
	updateRate := anqiTemplate.UpdateWfb
	blessMax := anqiTemplate.ZhufuMax
	addMin := anqiTemplate.AddMin
	addMax := anqiTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeAnqi)
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += anqiTemplate.AddMax
	}
	curTimesNum += addTimes

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//暗器祝福丹
func AnqiEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, anqiTemplate *gametemplate.AnqiTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeAnqi, anqiTemplate.TimesMin, anqiTemplate.TimesMax)
	updateRate := anqiTemplate.UpdateWfb
	blessMax := anqiTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//暗器丹培养判断
func AnqiDan(curTimesNum int32, curBless int32, anqiDanTemplate *gametemplate.AnqiDanTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := anqiDanTemplate.TimesMin
	timesMax := anqiDanTemplate.TimesMax
	updateRate := anqiDanTemplate.UpdateWfb
	blessMax := anqiDanTemplate.ZhufuMax
	addMin := anqiDanTemplate.AddMin
	addMax := anqiDanTemplate.AddMax + 1

	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//暗器进阶
func HandleAnqiAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeAnQiAdvanced)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("anqi:暗器升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	anqiManager := pl.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiInfo := anqiManager.GetAnqiInfo()
	beforeNum := int32(anqiInfo.AdvanceId)
	nextAdvancedId := anqiInfo.AdvanceId + 1
	anqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(nextAdvancedId))
	if anqiTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("anqi:暗器已达最高阶")

		playerlogic.SendSystemMessage(pl, lang.AnqiAdanvacedReachedLimit)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//进阶需要消耗的元宝
	costGold := anqiTemplate.UseMoney
	//进阶需要消耗的银两
	costSilver := int64(anqiTemplate.UseYinliang)
	//进阶需要消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := anqiTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := anqiTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = anqiTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("anqi:暗器进阶丹不足,无法进阶")
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
		// 	}).Warn("anqi:商铺没有该道具,无法自动购买")
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
				}).Warn("anqi:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("anqi:购买物品失败,自动进阶已停止")
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
			}).Warn("anqi:银两不足,无法进阶")
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
			}).Warn("anqi:元宝不足,无法进阶")
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
	goldUseReason := commonlog.GoldLogReasonAnqiAdvanced
	silverUseReason := commonlog.SilverLogReasonAnqiAdvanced
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), goldUseReason, goldUseReason.String(), costSilver, silverUseReason, silverUseReason.String())
	if !flag {
		panic(fmt.Errorf("anqi: anqiAdvanced Cost should be ok"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonAnqiAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), anqiInfo.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("anqi: anqiAdvanced use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//消耗事件
	if useItemTemplate != nil {
		gameevent.Emit(anqieventtypes.EventTypeAnqiAdvancedCost, pl, int32(anqiInfo.AdvanceId))
	}

	//进阶判断
	sucess, pro, randBless, addTimes, isDouble := AnqiAdvanced(pl, anqiInfo.TimesNum, anqiInfo.Bless, anqiTemplate)
	anqiManager.AnqiAdvanced(pro, addTimes, sucess)
	if sucess {
		//同步属性
		AnqiPropertyChanged(pl)
		advancedFinish(pl, int32(anqiInfo.AdvanceId))

		anqiReason := commonlog.AnqiLogReasonAdvanced
		reasonText := fmt.Sprintf(anqiReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := anqieventtypes.CreatePlayerAnqiAdvancedLogEventData(beforeNum, 1, anqiReason, reasonText)
		gameevent.Emit(anqieventtypes.EventTypeAnqiAdvancedLog, pl, data)
	} else {
		//进阶不成功
		advancedBless(pl, anqiInfo.AdvanceId, randBless, anqiInfo.Bless, anqiInfo.BlessTime, isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32) {
	scAnqiAdvanced := pbutil.BuildSCAnqiAdavancedFinshed(int32(advancedId), commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scAnqiAdvanced)
	return
}

func advancedBless(pl player.Player, advancedId int, bless int32, totalBless int32, blessTime int64, isDouble bool) {
	scAnqiAdvanced := pbutil.BuildSCAnqiAdavanced(int32(advancedId), bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scAnqiAdvanced)
	return
}
