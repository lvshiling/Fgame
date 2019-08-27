package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commomlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	"fgame/fgame/game/shihunfan/pbutil"
	"fgame/fgame/game/shop/shop"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	funcopentypes "fgame/fgame/game/funcopen/types"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	shoplogic "fgame/fgame/game/shop/logic"

	log "github.com/Sirupsen/logrus"
)

//噬魂幡进阶
func HandleShiHunFanAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeShiHunFanAdvanced)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shihunfan:噬魂幡升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	shihunfanManager := pl.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shihunfanInfo := shihunfanManager.GetShiHunFanInfo()
	beforeNum := int32(shihunfanInfo.AdvanceId)
	nextAdvancedId := beforeNum + 1
	nextTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(nextAdvancedId)
	if nextTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("shihunfan:噬魂幡已达最高阶")

		playerlogic.SendSystemMessage(pl, lang.ShiHunFanAdanvacedReachedLimit)
		return
	}

	if nextTemplate.GetAdvancedType() != commontypes.SpecialAdvancedTypeDefault {
		unusualAdvanced(pl, nextTemplate.GetAdvancedType(), nextTemplate.ShengJieValue1, shihunfanInfo.ChargeVal)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//进阶需要消耗的元宝
	costGold := nextTemplate.UseMoney
	//进阶需要消耗的银两
	costSilver := int64(nextTemplate.UseYinliang)
	//进阶需要消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := nextTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := nextTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = nextTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("shihunfan:噬魂幡进阶丹不足,无法进阶")
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
				}).Warn("shihunfan:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("shihunfan:购买物品失败,自动进阶已停止")
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
			}).Warn("shihunfan:银两不足,无法进阶")
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
			}).Warn("shihunfan:元宝不足,无法进阶")
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
			}).Warn("shihunfan:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonShiHunFanAdvanced
	silverUseReason := commonlog.SilverLogReasonShiHunFanAdvanced
	goldUseReasonText := fmt.Sprintf(commonlog.GoldLogReasonShiHunFanAdvanced.String(), shihunfanInfo.AdvanceId)
	silverUseReasonText := fmt.Sprintf(commonlog.SilverLogReasonShiHunFanAdvanced.String(), shihunfanInfo.AdvanceId)
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), goldUseReason, goldUseReasonText, costSilver, silverUseReason, silverUseReasonText)
	if !flag {
		panic(fmt.Errorf("shihunfan: shihunfanAdvanced Cost should be ok"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonShiHunFanAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), shihunfanInfo.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("shihunfan: shihunfanAdvanced use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	if useItemTemplate != nil {
		//消耗事件
		gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanAdvancedCost, pl, int32(shihunfanInfo.AdvanceId))
	}

	//进阶判断
	sucess, pro, _, addTimes, isDouble := ShiHunFanAdvanced(pl, shihunfanInfo.TimesNum, shihunfanInfo.Bless, nextTemplate)
	shihunfanManager.ShiHunFanAdvanced(pro, addTimes, sucess)
	if sucess {
		//同步属性
		ShiHunFanPropertyChanged(pl)

		shihunfanReason := commonlog.ShiHunFanLogReasonAdvanced
		reasonText := fmt.Sprintf(shihunfanReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := shihunfaneventtypes.CreatePlayerShiHunFanAdvancedLogEventData(beforeNum, 1, shihunfanReason, reasonText)
		gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanAdvancedLog, pl, data)
		advancedFinish(pl, shihunfanInfo)
	} else {
		//进阶不成功
		advancedBless(pl, shihunfanInfo, isDouble, autoFlag)
	}
	return
}

func unusualAdvanced(pl player.Player, advancedType commontypes.SpecialAdvancedType, needValue int32, nowVal int32) {
	switch advancedType {
	case commontypes.SpecialAdvancedTypeCharge:
		//充值数类型
		if needValue > nowVal {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needGold": needValue,
					"nowVal":   nowVal,
				}).Warn("shihunfan:噬魂幡升阶请求，不满足充值条件")
			playerlogic.SendSystemMessage(pl, lang.ShiHunFanAdvancedNotCharge)
			return
		}
		break
	default:
		return
	}

	shihunfanManager := pl.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shihunfanInfo := shihunfanManager.GetShiHunFanInfo()
	beforeNum := int32(shihunfanInfo.AdvanceId)

	shihunfanManager.ShiHunFanAdvanced(0, 0, true)
	//同步属性
	ShiHunFanPropertyChanged(pl)

	shihunfanReason := commonlog.ShiHunFanLogReasonAdvanced
	reasonText := fmt.Sprintf(shihunfanReason.String(), commontypes.SpecialAdvancedTypeCharge.String())
	data := shihunfaneventtypes.CreatePlayerShiHunFanAdvancedLogEventData(beforeNum, 1, shihunfanReason, reasonText)
	gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanAdvancedLog, pl, data)
	advancedFinish(pl, shihunfanInfo)
	return
}

func advancedFinish(pl player.Player, info *playershihunfan.PlayerShiHunFanObject) {
	scShiHunFanAdvanced := pbutil.BuildSCShiHunFanAdavancedFinshed(info, commontypes.AdvancedTypeXingChen)
	pl.SendMsg(scShiHunFanAdvanced)
	return
}

func advancedBless(pl player.Player, info *playershihunfan.PlayerShiHunFanObject, isDouble bool, isAutoBuy bool) {
	scShiHunFanAdvanced := pbutil.BuildSCShiHunFanAdavanced(info, commontypes.AdvancedTypeXingChen, isDouble, isAutoBuy)
	pl.SendMsg(scShiHunFanAdvanced)
	return
}

//变更噬魂幡系统属性
func ShiHunFanPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeShiHunFan.Mask())
	return
}

//噬魂幡系统进阶判断
func ShiHunFanAdvanced(pl player.Player, curTimesNum int32, curBless int32, shiHunFanTemplate *gametemplate.ShiHunFanTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeShiHunFan, shiHunFanTemplate.TimesMin, shiHunFanTemplate.TimesMax)
	updateRate := shiHunFanTemplate.UpdateWfb
	blessMax := shiHunFanTemplate.ZhufuMax
	addMin := shiHunFanTemplate.AddMin
	addMax := shiHunFanTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeShiHunFan)
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += shiHunFanTemplate.AddMax
	}

	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//噬魂幡祝福丹
func ShiHunFanEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, shiHunFanTemplate *gametemplate.ShiHunFanTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeAnqi, shiHunFanTemplate.TimesMin, shiHunFanTemplate.TimesMax)
	updateRate := shiHunFanTemplate.UpdateWfb
	blessMax := shiHunFanTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}
