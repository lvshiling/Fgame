package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commomlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"

	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotemplate "fgame/fgame/game/fabao/template"
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

//法宝进阶的逻辑
func HandleFaBaoAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeFaBaoAdvanced)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("fabao:法宝升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoInfo := manager.GetFaBaoInfo()
	beforeNum := faBaoInfo.GetAdvancedId()
	nextAdvancedId := beforeNum + 1
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(nextAdvancedId))
	if faBaoTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Warn("FaBao:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.FaBaoAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := faBaoTemplate.UseMoney
	//进阶需要消耗的银两
	costSilver := int64(faBaoTemplate.UseYinliang)
	//进阶需要消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := faBaoTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := faBaoTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = faBaoTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("FaBao:法宝进阶丹不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动进阶
		needBuyNum := itemCount - totalNum
		itemCount = totalNum

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("fabao:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("fabao:购买物品失败,自动进阶已停止")
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
			}).Warn("fabao:银两不足,无法进阶")
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
			}).Warn("fabao:元宝不足,无法进阶")
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
			}).Warn("fabao:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonFaBaoAdvanced.String()
	reasonSliverText := commonlog.SilverLogReasonFaBaoAdvanced.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonFaBaoAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonFaBaoAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("fabao: faBaoAdvanced Cost should be ok"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonFaBaoAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), faBaoInfo.GetAdvancedId())
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("fabao: faBaoAdvanced use item shoud be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	if useItemTemplate != nil {
		//消耗事件
		gameevent.Emit(fabaoeventtypes.EventTypeFaBaoAdvancedCost, pl, faBaoInfo.GetAdvancedId())
	}

	//法宝进阶判断
	sucess, pro, randBless, addTimes, isDouble := FaBaoAdvanced(pl, faBaoInfo.GetTimesNum(), faBaoInfo.GetBless(), faBaoTemplate)
	manager.FaBaoAdvanced(pro, addTimes, sucess)
	if sucess {
		faBaoReason := commonlog.FaBaoLogReasonAdvanced
		reasonText := fmt.Sprintf(faBaoReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := fabaoeventtypes.CreatePlayerFaBaoAdvancedLogEventData(beforeNum, 1, faBaoReason, reasonText)
		gameevent.Emit(fabaoeventtypes.EventTypeFaBaoAdvancedLog, pl, data)

		//同步属性
		FaBaoPropertyChanged(pl)
		advancedFinish(pl, int32(faBaoInfo.GetAdvancedId()), faBaoInfo.GetFaBaoId())
	} else {
		//进阶不成功
		advancedBless(pl, faBaoInfo.GetAdvancedId(), faBaoInfo.GetFaBaoId(), randBless, faBaoInfo.GetBless(), faBaoInfo.GetBlessTime(), isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32, faBaoId int32) (err error) {
	scFaBaoAdvanced := pbutil.BuildSCFaBaoAdavancedFinshed(advancedId, faBaoId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scFaBaoAdvanced)
	return
}

func advancedBless(pl player.Player, advancedId int32, faBaoId int32, bless int32, totalBless int32, blessTime int64, isDouble bool) (err error) {
	scFaBaoAdvanced := pbutil.BuildSCFaBaoAdavanced(advancedId, faBaoId, bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scFaBaoAdvanced)
	return
}

//变更法宝属性
func FaBaoPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeFaBao.Mask())
	return
}

//法宝进阶判断
func FaBaoAdvanced(pl player.Player, curTimesNum int32, curBless int32, faBaoTemplate *gametemplate.FaBaoTemplate) (sucess bool, pro int32, randBless int32, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeFaBao, faBaoTemplate.TimesMin, faBaoTemplate.TimesMax)
	updateRate := faBaoTemplate.UpdateWfb
	blessMax := faBaoTemplate.ZhufuMax
	addMin := faBaoTemplate.AddMin
	addMax := faBaoTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeFaBao)
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += faBaoTemplate.AddMax
	}

	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//法宝皮肤升星判断
func FaBaoSkinUpstar(curTimesNum int32, curBless int32, upstarTemplate *gametemplate.FaBaoUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upstarTemplate.TimesMin
	timesMax := upstarTemplate.TimesMax
	updateRate := upstarTemplate.UpstarRate
	blessMax := upstarTemplate.ZhufuMax
	addMin := upstarTemplate.AddMin
	addMax := upstarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//法宝通灵
func FaBaoTongLing(curTimesNum int32, curBless int32, tongLingTemplate *gametemplate.FaBaoTongLingTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := tongLingTemplate.TimesMin
	timesMax := tongLingTemplate.TimesMax
	updateRate := tongLingTemplate.UpdateWfb
	blessMax := tongLingTemplate.ZhufuMax
	addMin := tongLingTemplate.AddMin
	addMax := tongLingTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//法宝祝福丹
func FaBaoEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, fabaoTemplate *gametemplate.FaBaoTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeFaBao, fabaoTemplate.TimesMin, fabaoTemplate.TimesMax)
	updateRate := fabaoTemplate.UpdateWfb
	blessMax := fabaoTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}
