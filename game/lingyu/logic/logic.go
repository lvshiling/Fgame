package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commomlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"

	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
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

//领域进阶的逻辑
func HandleLingyuAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeLingYuAdvanced)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lingyu:领域升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()
	beforeNum := int32(lingyuInfo.AdvanceId)
	nextAdvancedId := int32(lingyuInfo.AdvanceId + 1)
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(nextAdvancedId)
	if lingyuTemplate == nil {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("Lingyu:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.LingyuAdanvacedReachedLimit)
		return
	}
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//消费激活
	if lingyuTemplate.GetActivateType() == commontypes.SpecialAdvancedTypeCost {
		needActivateGold := int64(lingyuTemplate.ShengjieValue)
		if !propertyManager.HasEnoughGold(needActivateGold, false) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needGold": lingyuTemplate.ShengjieValue,
					"nowVal":   lingyuInfo.ChargeVal,
				}).Warn("lingyu:领域升阶请求，元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}

		useGoldReason := commonlog.GoldLogReasonLingyuActivate
		flag := propertyManager.CostGold(needActivateGold, false, useGoldReason, useGoldReason.String())
		if !flag {
			panic("lingyu:激活消耗应该成功")
		}
		propertylogic.SnapChangedProperty(pl)

	} else {
		//进阶需要消耗的元宝
		costGold := lingyuTemplate.UseMoney
		//进阶需要消耗的银两
		costSilver := int64(lingyuTemplate.UseYinliang)
		//进阶需要消耗绑元
		costBindGold := int32(0)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := lingyuTemplate.UseItem
		isEnoughBuyTimes := true
		shopIdMap := make(map[int32]int32)
		useItemTemplate := lingyuTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = lingyuTemplate.ItemCount
			totalNum = inventoryManager.NumOfItems(int32(useItem))
		}
		if totalNum < itemCount {
			if autoFlag == false {
				log.WithFields(log.Fields{
					"playerid": pl.GetId(),
				}).Warn("Lingyu:领域进阶丹不足,无法进阶")
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
			// 		"playerid": pl.GetId(),
			// 	}).Warn("lingyu:商铺没有该道具,无法自动购买")
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
					}).Warn("lingyu:商铺没有该道具,无法自动购买")
					playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
					return
				}

				isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
				if !isEnoughBuyTimes {
					log.WithFields(log.Fields{
						"playerId": pl.GetId(),
						"autoFlag": autoFlag,
					}).Warn("lingyu:购买物品失败,自动进阶已停止")
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
					"playerid": pl.GetId(),
				}).Warn("lingyu:银两不足,无法进阶")
				playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
				return
			}
		}
		//是否足够元宝
		if costGold != 0 {
			flag := propertyManager.HasEnoughGold(int64(costGold), false)
			if !flag {
				log.WithFields(log.Fields{
					"playerid": pl.GetId(),
				}).Warn("lingyu:元宝不足,无法进阶")
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
					"playerid": pl.GetId(),
				}).Warn("lingyu:元宝不足,无法进阶")
				playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
				return
			}
		}

		//更新自动购买每日限购次数
		if len(shopIdMap) != 0 {
			shoplogic.ShopDayCountChanged(pl, shopIdMap)
		}

		//消耗钱
		reasonGoldText := commonlog.GoldLogReasonLingyuAdvanced.String()
		reasonSliverText := commonlog.SilverLogReasonLingyuAdvanced.String()
		flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonLingyuAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonLingyuAdvanced, reasonSliverText)
		if !flag {
			panic(fmt.Errorf("lingyu: lingyuAdvanced Cost should be ok"))
		}
		//同步元宝
		if costGold != 0 || costSilver != 0 || costBindGold != 0 {
			propertylogic.SnapChangedProperty(pl)
		}

		//消耗物品
		if itemCount != 0 {
			inventoryReason := commonlog.InventoryLogReasonLingyuAdvanced
			reasonText := fmt.Sprintf(inventoryReason.String(), lingyuInfo.AdvanceId)
			flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("lingyu: lingyuAdvanced use item shoud be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)

		}
		if useItemTemplate != nil {
			//消耗事件
			gameevent.Emit(lingyueventtypes.EventTypeLingyuAdvancedCost, pl, int32(lingyuInfo.AdvanceId))
		}
	}

	//领域进阶判断
	sucess, pro, randBless, addTimes, isDouble := LingYuAdvanced(pl, lingyuInfo.TimesNum, lingyuInfo.Bless, lingyuTemplate)
	lingyuManager.LingyuAdvanced(pro, addTimes, sucess)
	if sucess {
		//同步属性
		LingyuPropertyChanged(pl)
		advancedFinish(pl, int32(lingyuInfo.AdvanceId), lingyuInfo.LingyuId)

		lingyuReason := commonlog.LingyuLogReasonAdvanced
		reasonText := fmt.Sprintf(lingyuReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := lingyueventtypes.CreatePlayerLingyuAdvancedLogEventData(beforeNum, 1, lingyuReason, reasonText)
		gameevent.Emit(lingyueventtypes.EventTypeLingyuAdvancedLog, pl, data)
	} else {
		//进阶不成功
		advancedBless(pl, lingyuInfo.AdvanceId, lingyuInfo.LingyuId, randBless, lingyuInfo.Bless, lingyuInfo.BlessTime, isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32, lingyuId int32) {
	scLingyuAdvanced := pbutil.BuildSCLingyuAdavancedFinshed(int32(advancedId), lingyuId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scLingyuAdvanced)

}

func advancedBless(pl player.Player, advancedId int, lingyuId int32, bless int32, totalBless int32, blessTime int64, isDouble bool) {
	scLingyuAdvanced := pbutil.BuildSCLingyuAdavanced(int32(advancedId), lingyuId, bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scLingyuAdvanced)

}

//变更领域属性
func LingyuPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeLingyu.Mask())
	return
}

//领域进阶判断
func LingYuAdvanced(pl player.Player, curTimesNum int32, curBless int32, lingYuTemplate *gametemplate.LingyuTemplate) (sucess bool, pro int32, randBless, addTimes int32, isDouble bool) {
	if lingYuTemplate.GetActivateType() != commontypes.SpecialAdvancedTypeDefault {
		sucess = true
		return
	}

	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeLingyu, lingYuTemplate.TimesMin, lingYuTemplate.TimesMax)
	updateRate := lingYuTemplate.UpdateWfb
	blessMax := lingYuTemplate.ZhufuMax
	addMin := lingYuTemplate.AddMin
	addMax := lingYuTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeLingyu)
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += lingYuTemplate.AddMax
	}
	curTimesNum += addTimes

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//领域祝福丹判断
func LingEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, lingYuTemplate *gametemplate.LingyuTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeLingyu, lingYuTemplate.TimesMin, lingYuTemplate.TimesMax)
	updateRate := lingYuTemplate.UpdateWfb
	blessMax := lingYuTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//领域幻化丹培养判断
func LingYuHuanHua(curTimesNum int32, curBless int32, huanHuaTemplate *gametemplate.LingyuHuanHuaTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := huanHuaTemplate.TimesMin
	timesMax := huanHuaTemplate.TimesMax
	updateRate := huanHuaTemplate.UpdateWfb
	blessMax := huanHuaTemplate.ZhufuMax
	addMin := huanHuaTemplate.AddMin
	addMax := huanHuaTemplate.AddMax + 1

	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//领域皮肤升星判断
func LingYuSkinUpstar(curTimesNum int32, curBless int32, upstarTemplate *gametemplate.FieldUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upstarTemplate.TimesMin
	timesMax := upstarTemplate.TimesMax
	updateRate := upstarTemplate.UpstarRate
	blessMax := upstarTemplate.ZhufuMax
	addMin := upstarTemplate.AddMin
	addMax := upstarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
