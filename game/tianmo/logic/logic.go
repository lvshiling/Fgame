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
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	gametemplate "fgame/fgame/game/template"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	"fgame/fgame/game/tianmo/pbutil"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotemplate "fgame/fgame/game/tianmo/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//变更天魔属性
func TianMoPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeTianMoTi.Mask())
	return
}

//天魔进阶判断
func TianMoAdvanced(pl player.Player, curTimesNum int32, curBless int32, tianMoTemplate *gametemplate.TianMoTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	if tianMoTemplate.GetActivateType() == commontypes.SpecialAdvancedTypeCharge {
		sucess = true
		return
	}

	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeTianMoTi, tianMoTemplate.TimesMin, tianMoTemplate.TimesMax)
	updateRate := tianMoTemplate.UpdateWfb
	blessMax := tianMoTemplate.ZhufuMax
	addMin := tianMoTemplate.AddMin
	addMax := tianMoTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeTianMoTi)
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += tianMoTemplate.AddMax
	}
	curTimesNum += addTimes

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//天魔祝福丹
func TianMoEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, tianMoTemplate *gametemplate.TianMoTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeTianMoTi, tianMoTemplate.TimesMin, tianMoTemplate.TimesMax)
	updateRate := tianMoTemplate.UpdateWfb
	blessMax := tianMoTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//天魔丹培养判断
func TianMoDan(curTimesNum int32, curBless int32, danTemplate *gametemplate.TianMoDanTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := danTemplate.TimesMin
	timesMax := danTemplate.TimesMax
	updateRate := danTemplate.UpdateWfb
	blessMax := danTemplate.ZhufuMax
	addMin := danTemplate.AddMin
	addMax := danTemplate.AddMax + 1

	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//天魔进阶
func HandleTianMoAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeTianMoAdvanced)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tianmo:天魔体升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	tianmoManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianmoInfo := tianmoManager.GetTianMoInfo()
	beforeNum := tianmoInfo.AdvanceId
	nextAdvancedId := tianmoInfo.AdvanceId + 1
	tianMoTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(int32(nextAdvancedId))
	if tianMoTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("tianmo:天魔已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.TianMoAdanvacedReachedLimit)
		return
	}

	//充值激活
	if tianMoTemplate.GetActivateType() == commontypes.SpecialAdvancedTypeCharge {
		if tianMoTemplate.ShengjieValue > int32(tianmoInfo.ChargeVal) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needGold": tianMoTemplate.ShengjieValue,
					"nowVal":   tianmoInfo.ChargeVal,
				}).Warn("tianmo:天魔体升阶请求，不满足充值条件")
			playerlogic.SendSystemMessage(pl, lang.TianMoActivateNotEnoughCharge)
			return
		}
	} else {
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

		//进阶需要消耗的元宝
		costGold := tianMoTemplate.UseMoney
		//进阶需要消耗的银两
		costSilver := int64(tianMoTemplate.UseYinliang)
		//进阶需要消耗的绑元
		costBindGold := int32(0)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := tianMoTemplate.UseItem
		isEnoughBuyTimes := true
		shopIdMap := make(map[int32]int32)
		useItemTemplate := tianMoTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = tianMoTemplate.ItemCount
			totalNum = inventoryManager.NumOfItems(int32(useItem))
		}
		if totalNum < itemCount {
			if autoFlag == false {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("tianMo:天魔进阶丹不足,无法进阶")
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
			// 	}).Warn("tianMo:商铺没有该道具,无法自动购买")
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
					}).Warn("tianMo:商铺没有该道具,无法自动购买")
					playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
					return
				}

				isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
				if !isEnoughBuyTimes {
					log.WithFields(log.Fields{
						"playerId": pl.GetId(),
						"autoFlag": autoFlag,
					}).Warn("tianMo:购买物品失败,自动进阶已停止")
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
				}).Warn("tianMo:银两不足,无法进阶")
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
				}).Warn("tianMo:元宝不足,无法进阶")
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
		goldUseReason := commonlog.GoldLogReasonTianMoAdvanced
		silverUseReason := commonlog.SilverLogReasonTianMoAdvanced
		flag := propertyManager.Cost(int64(costBindGold), int64(costGold), goldUseReason, goldUseReason.String(), costSilver, silverUseReason, silverUseReason.String())
		if !flag {
			panic(fmt.Errorf("tianMo: tianMoAdvanced Cost should be ok"))
		}
		//同步元宝
		if costGold != 0 || costSilver != 0 || costBindGold != 0 {
			propertylogic.SnapChangedProperty(pl)
		}

		//消耗物品
		if itemCount != 0 {
			inventoryReason := commonlog.InventoryLogReasonTianMoAdvanced
			reasonText := fmt.Sprintf(inventoryReason.String(), tianmoInfo.AdvanceId)
			flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("tianmo: tianmoAdvanced use item should be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)
		}

		//消耗事件
		if useItemTemplate != nil {
			gameevent.Emit(tianmoeventtypes.EventTypeTianMoAdvancedCost, pl, int32(tianmoInfo.AdvanceId))
		}
	}

	//进阶判断
	sucess, pro, randBless, addTimes, isDouble := TianMoAdvanced(pl, tianmoInfo.TimesNum, tianmoInfo.Bless, tianMoTemplate)
	tianmoManager.TianMoAdvanced(pro, addTimes, sucess)
	if sucess {
		//同步属性
		TianMoPropertyChanged(pl)
		advancedFinish(pl, int32(tianmoInfo.AdvanceId))

		tianmoReason := commonlog.TianMoLogReasonAdvanced
		reasonText := fmt.Sprintf(tianmoReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := tianmoeventtypes.CreatePlayerTianMoAdvancedLogEventData(beforeNum, 1, tianmoReason, reasonText)
		gameevent.Emit(tianmoeventtypes.EventTypeTianMoAdvancedLog, pl, data)
	} else {
		//进阶不成功
		advancedBless(pl, tianmoInfo.AdvanceId, randBless, tianmoInfo.Bless, tianmoInfo.BlessTime, isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32) {
	scTianMoAdvanced := pbutil.BuildSCTianMoAdavancedFinshed(int32(advancedId), commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scTianMoAdvanced)
	return
}

func advancedBless(pl player.Player, advancedId int32, bless int32, totalBless int32, blessTime int64, isDouble bool) {
	scTianMoAdvanced := pbutil.BuildSCTianMoAdavanced(advancedId, bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scTianMoAdvanced)
	return
}
