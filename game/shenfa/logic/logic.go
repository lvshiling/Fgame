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
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"

	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
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

//身法进阶的逻辑
func HandleShenfaAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenfaAdvanced)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenfa:身法升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfaInfo := shenfaManager.GetShenfaInfo()
	nextAdvancedId := int32(shenfaInfo.AdvanceId + 1)
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(nextAdvancedId)
	if shenfaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("Shenfa:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.ShenfaAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的绑元
	costGold := shenfaTemplate.UseMoney
	//进阶需要消耗的银两
	costSilver := int64(shenfaTemplate.UseYinliang)
	//进阶需要消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := shenfaTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := shenfaTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = shenfaTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}

	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(
				log.Fields{
					"playerid": pl.GetId(),
				}).Warn("Shenfa:身法进阶丹不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动进阶
		needBuyNum := itemCount - totalNum
		itemCount = totalNum
		//获取价格
		// shopTemplate := shop.GetShopService().GetShopTemplateByItem(useItem)
		// if shopTemplate == nil {
		// 	log.WithFields(
		// 		log.Fields{
		// 			"playerid": pl.GetId(),
		// 		}).Warn("shenfa:商铺没有该道具,无法自动购买")
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
				}).Warn("shenfa:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("shenfa:购买物品失败,自动进阶已停止")
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
			log.WithFields(
				log.Fields{
					"playerid": pl.GetId(),
				}).Warn("shenfa:银两不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerid": pl.GetId(),
				}).Warn("shenfa:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerid": pl.GetId(),
				}).Warn("shenfa:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonShenfaAdvanced.String()
	reasonSliverText := commonlog.SilverLogReasonShenfaAdvanced.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonShenfaAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonShenfaAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("shenfa: shenfaAdvanced Cost should be ok"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonShenfaAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), shenfaInfo.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("shenfa: shenfaAdvanced use item shoud be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	if useItemTemplate != nil {
		//消耗事件
		gameevent.Emit(shenfaeventtypes.EventTypeShenfaAdvancedCost, pl, int32(shenfaInfo.AdvanceId))
	}

	//身法进阶判断
	beforeNum := int32(shenfaInfo.AdvanceId)
	sucess, pro, randBless, addTimes, isDouble := ShenFaAdvanced(pl, shenfaInfo.TimesNum, shenfaInfo.Bless, shenfaTemplate)
	shenfaManager.ShenfaAdvanced(pro, addTimes, sucess)
	if sucess {
		shenfaReason := commonlog.ShenfaLogReasonAdvanced
		reasonText := fmt.Sprintf(shenfaReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := shenfaeventtypes.CreatePlayerShenfaAdvancedLogEventData(beforeNum, 1, shenfaReason, reasonText)
		gameevent.Emit(shenfaeventtypes.EventTypeShenfaAdvancedLog, pl, data)

		//同步属性
		ShenfaPropertyChanged(pl)
		advancedFinish(pl, int32(shenfaInfo.AdvanceId), shenfaInfo.ShenfaId)
	} else {
		//进阶不成功
		advancedBless(pl, shenfaInfo.AdvanceId, shenfaInfo.ShenfaId, randBless, shenfaInfo.Bless, shenfaInfo.BlessTime, isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32, shenfaId int32) {
	scShenfaAdvanced := pbutil.BuildSCShenfaAdavancedFinshed(int32(advancedId), shenfaId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scShenfaAdvanced)

}

func advancedBless(pl player.Player, advancedId int, shenfaId int32, bless int32, totalBless int32, blessTime int64, isDouble bool) {
	scShenfaAdvanced := pbutil.BuildSCShenfaAdavanced(int32(advancedId), shenfaId, bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scShenfaAdvanced)

}

//变更身法属性
func ShenfaPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeShenfa.Mask())

	return
}

//身法进阶判断
func ShenFaAdvanced(pl player.Player, curTimesNum int32, curBless int32, shenFaTemplate *gametemplate.ShenfaTemplate) (sucess bool, pro int32, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeShenfa, shenFaTemplate.TimesMin, shenFaTemplate.TimesMax)
	updateRate := shenFaTemplate.UpdateWfb
	blessMax := shenFaTemplate.ZhufuMax
	addMin := shenFaTemplate.AddMin
	addMax := shenFaTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeShenfa)
	if isDouble {
		addTimes += extralAddTimes
		randBless += shenFaTemplate.AddMax
	}
	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//身法祝福丹判断
func ShenFaEatZhuFuDan(pl player.Player, curTimesNum int32, curBless, randBless int32, shenFaTemplate *gametemplate.ShenfaTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeShenfa, shenFaTemplate.TimesMin, shenFaTemplate.TimesMax)
	updateRate := shenFaTemplate.UpdateWfb
	blessMax := shenFaTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//身法幻化丹培养判断
func ShenFaHuanHua(curTimesNum int32, curBless int32, huanHuaTemplate *gametemplate.ShenfaHuanHuaTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := huanHuaTemplate.TimesMin
	timesMax := huanHuaTemplate.TimesMax
	updateRate := huanHuaTemplate.UpdateWfb
	blessMax := huanHuaTemplate.ZhufuMax
	addMin := huanHuaTemplate.AddMin
	addMax := huanHuaTemplate.AddMax + 1

	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//身法皮肤升星判断
func ShenFaSkinUpstar(curTimesNum int32, curBless int32, upstarTemplate *gametemplate.ShenFaUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upstarTemplate.TimesMin
	timesMax := upstarTemplate.TimesMax
	updateRate := upstarTemplate.UpstarRate
	blessMax := upstarTemplate.ZhufuMax
	addMin := upstarTemplate.AddMin
	addMax := upstarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
