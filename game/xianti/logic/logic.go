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
	xiantieventtypes "fgame/fgame/game/xianti/event/types"

	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
	"fgame/fgame/game/xianti/xianti"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//仙体进阶的逻辑
func HandleXianTiAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeXianTiAdvanced)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xianti:仙体升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	xianTiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiInfo := xianTiManager.GetXianTiInfo()
	beforeNum := int32(xianTiInfo.AdvanceId)
	nextAdvancedId := xianTiInfo.AdvanceId + 1
	xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(nextAdvancedId))
	if xianTiTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Warn("xianti:仙体已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.XianTiAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := xianTiTemplate.UseMoney
	//进阶需要消耗的银两
	costSilver := int64(xianTiTemplate.UseYinliang)
	//进阶需要的消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := xianTiTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := xianTiTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = xianTiTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("xianti:仙体进阶丹不足,无法进阶")
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
				}).Warn("xianti:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("xianti:购买物品失败,自动进阶已停止")
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
			}).Warn("xianti:银两不足,无法进阶")
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
			}).Warn("xianti:元宝不足,无法进阶")
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
			}).Warn("xianti:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonXianTiAdvanced.String()
	reasonSliverText := commonlog.SilverLogReasonXianTiAdvanced.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonXianTiAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonXianTiAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("xianti: xianTiAdvanced Cost should be ok"))
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonXianTiAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), xianTiInfo.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("xianti: xianTiAdvanced use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	if useItemTemplate != nil {
		//消耗事件
		gameevent.Emit(xiantieventtypes.EventTypeXianTiAdvancedCost, pl, int32(xianTiInfo.AdvanceId))
	}

	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//进阶判断
	sucess, pro, randBless, addTimes, isDouble := XianTiAdvanced(pl, xianTiInfo.TimesNum, xianTiInfo.Bless, xianTiTemplate)
	xianTiManager.XianTiAdvanced(pro, addTimes, sucess)
	if sucess {
		xianTiReason := commonlog.XianTiLogReasonAdvanced
		reasonText := fmt.Sprintf(xianTiReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := xiantieventtypes.CreatePlayerXianTiAdvancedLogEventData(beforeNum, 1, xianTiReason, reasonText)
		gameevent.Emit(xiantieventtypes.EventTypeXianTiAdvancedLog, pl, data)

		//同步属性
		XianTiPropertyChanged(pl)
		advancedFinish(pl, int32(xianTiInfo.AdvanceId), xianTiInfo.XianTiId)
	} else {
		//进阶不成功
		advancedBless(pl, xianTiInfo.AdvanceId, xianTiInfo.XianTiId, randBless, xianTiInfo.Bless, xianTiInfo.BlessTime, isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32, xianTiId int32) (err error) {
	scXianTiAdvanced := pbutil.BuildSCXianTiAdavancedFinshed(advancedId, xianTiId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scXianTiAdvanced)
	return
}

func advancedBless(pl player.Player, advancedId int, xianTiId int32, bless int32, totalBless int32, blessTime int64, isDouble bool) (err error) {
	scXianTiAdvanced := pbutil.BuildSCXianTiAdavanced(int32(advancedId), xianTiId, bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scXianTiAdvanced)
	return
}

//变更仙体属性
func XianTiPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeXianTi.Mask())
	return
}

//仙体进阶判断
func XianTiAdvanced(pl player.Player, curTimesNum int32, curBless int32, XianTiTemplate *gametemplate.XianTiTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeXianTi, XianTiTemplate.TimesMin, XianTiTemplate.TimesMax)
	updateRate := XianTiTemplate.UpdateWfb
	blessMax := XianTiTemplate.ZhufuMax
	addMin := XianTiTemplate.AddMin
	addMax := XianTiTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeXianTi)
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += XianTiTemplate.AddMax
	}
	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//仙体祝福丹判断
func XianTiEatZhuFuDan(pl player.Player, curTimesNum, curBless, randBless int32, XianTiTemplate *gametemplate.XianTiTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeXianTi, XianTiTemplate.TimesMin, XianTiTemplate.TimesMax)
	updateRate := XianTiTemplate.UpdateWfb
	blessMax := XianTiTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//仙体幻化丹培养判断
func XianTiHunaHuaFeed(curTimesNum int32, curBless int32, huanHuaTemplate *gametemplate.XianTiHuanHuaTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := huanHuaTemplate.TimesMin
	timesMax := huanHuaTemplate.TimesMax
	updateRate := huanHuaTemplate.UpdateWfb
	blessMax := huanHuaTemplate.ZhufuMax
	addMin := huanHuaTemplate.AddMin
	addMax := huanHuaTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//仙体皮肤升星判断
func XianTiSkinUpstar(curTimesNum int32, curBless int32, upstarTemplate *gametemplate.XianTiUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upstarTemplate.TimesMin
	timesMax := upstarTemplate.TimesMax
	updateRate := upstarTemplate.UpstarRate
	blessMax := upstarTemplate.ZhufuMax
	addMin := upstarTemplate.AddMin
	addMax := upstarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
