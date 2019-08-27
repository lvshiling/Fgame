package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commomlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongdevevent "fgame/fgame/game/lingtongdev/event/types"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//灵童养成类进阶的逻辑
func HandleLingTongDevAdvanced(pl player.Player, classType types.LingTongDevSysType, autoFlag bool) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"autoFlag":  autoFlag,
		}).Warn("LingTongDev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	isFuncOpen := pl.IsFuncOpen(classType.GetAdvanceFuncOpenType())
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn(fmt.Sprintf("LingTongDev:%s升阶功能未开启", classType.String()))
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	lingTongManager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongMap := lingTongManager.GetLingTongMap()
	if len(lingTongMap) == 0 {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"autoFlag":  autoFlag,
		}).Warn("LingTongDev:您还未激活过灵童,请先激活一只灵童")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevNoActivateLingTong)
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	// if lingTongDevInfo == nil {
	// 	lingTongDevInfo = manager.AdvancedInit(classType)
	// }
	beforeNum := lingTongDevInfo.GetAdvancedId()
	nextAdvancedId := beforeNum + 1
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, nextAdvancedId)
	if lingTongDevTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"autoFlag":  autoFlag,
		}).Warn("LingTongDev:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevAdanvacedReachedLimit, classType.String())
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := lingTongDevTemplate.GetGold()
	//进阶需要消耗的银两
	costSilver := lingTongDevTemplate.GetSilver()
	//进阶需要消耗的绑元
	costBindGold := int64(0)

	needItems := lingTongDevTemplate.GetItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag && autoFlag == false {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"autoFlag":  autoFlag,
			}).Warn("LingTongDev:道具不足，无法进阶")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	//获取背包物品和需要购买物品
	items, buyItems := inventoryManager.GetItemsAndNeedBuy(needItems)
	//计算需要元宝等
	if len(buyItems) != 0 {
		isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayerMap(pl, buyItems)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"autoFlag":  autoFlag,
			}).Warn("LingTongDev:购买物品失败,自动升星已停止")
			playerlogic.SendSystemMessage(pl, lang.ShopUpstarAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		costGold += shopNeedGold
		costBindGold += shopNeedBindGold
		costSilver += shopNeedSilver
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(costSilver)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"autoFlag":  autoFlag,
			}).Warn("lingtongdev:银两不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(costGold, false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"autoFlag":  autoFlag,
			}).Warn("lingtongdev:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够元宝
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(needBindGold, true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"autoFlag":  autoFlag,
			}).Warn("lingtongdev:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	flag := propertyManager.HasEnoughCost(costBindGold, costGold, costSilver)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"autoFlag":  autoFlag,
		}).Warn("lingtongdev:元宝不足，无法进阶")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	reasonGoldText := fmt.Sprintf(commonlog.GoldLogReasonLingTongAdvanced.String(), classType.String())
	reasonSliverText := fmt.Sprintf(commonlog.SilverLogReasonLingTongAdvanced.String(), classType.String())
	flag = propertyManager.Cost(costBindGold, costGold, commonlog.GoldLogReasonLingTongAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonLingTongAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("lingtongdev: lingTongDevAdvanced Cost should be ok"))
	}
	propertylogic.SnapChangedProperty(pl)

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗物品
	if len(items) != 0 {
		inventoryReason := commonlog.InventoryLogReasonLingTongAdvaced
		reasonText := fmt.Sprintf(inventoryReason.String(), classType.String(), lingTongDevInfo.GetAdvancedId())
		flag := inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtongdev: lingTongDevAdvanced use item shoud be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	if len(needItems) != 0 {
		//消耗事件
		gameevent.Emit(lingtongdevevent.EventTypeLingTongDevAdvancedCost, pl, lingTongDevInfo)
	}

	//灵童养成类进阶判断
	sucess, pro, randBless, addTimes, isDouble := LingTongDevAdvanced(pl, lingTongDevInfo.GetTimesNum(), lingTongDevInfo.GetBless(), lingTongDevTemplate)
	manager.LingTongDevAdvanced(classType, pro, addTimes, sucess)
	if sucess {
		lingTongDevReason := commonlog.LingTongDevLogReasonAdvanced
		reasonText := fmt.Sprintf(lingTongDevReason.String(), classType.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := lingtongdeveventtypes.CreatePlayerLingTongDevAdvancedLogEventData(classType, beforeNum, 1, lingTongDevReason, reasonText)
		gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevAdvancedLog, pl, data)

		//同步属性
		LingTongDevPropertyChanged(pl, classType)
		advancedFinish(pl, int32(classType), int32(lingTongDevInfo.GetAdvancedId()), lingTongDevInfo.GetSeqId())
	} else {
		//进阶不成功
		advancedBless(pl, int32(classType), lingTongDevInfo.GetAdvancedId(), lingTongDevInfo.GetSeqId(), randBless, lingTongDevInfo.GetBless(), lingTongDevInfo.GetBlessTime(), isDouble)
	}
	return
}

func advancedFinish(pl player.Player, classType int32, advancedId int32, seqId int32) (err error) {
	scLingTongDevAdvanced := pbutil.BuildSCLingTongDevAdavancedFinshed(classType, advancedId, seqId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scLingTongDevAdvanced)
	return
}

func advancedBless(pl player.Player, classType int32, advancedId int32, seqId int32, bless int32, totalBless int32, blessTime int64, isDouble bool) (err error) {
	scLingTongDevAdvanced := pbutil.BuildSCLingTongDevAdavanced(classType, advancedId, seqId, bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scLingTongDevAdvanced)
	return
}

//变更灵童属性属性
func LingTongDevPropertyChanged(pl player.Player, classType types.LingTongDevSysType) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyEffectorType := classType.GetPlayerPropertyEffectType()
	propertyManager.UpdateBattleProperty(propertyEffectorType.Mask())
	// /更新灵童属性
	LingTongDevSelfPropertyChanged(pl, classType)
	//写注册器会重复引用
	// propertyEffectorType := playerpropertytypes.PlayerPropertyEffectorTypeLingTongFaBao
	// switch classType {
	// case types.LingTongDevSysTypeLingBing:
	// 	propertyEffectorType = playerpropertytypes.PlayerPropertyEffectorTypeLingTongWeapon
	// case types.LingTongDevSysTypeLingQi:
	// 	propertyEffectorType = playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount
	// case types.LingTongDevSysTypeLingShen:
	// 	propertyEffectorType = playerpropertytypes.PlayerPropertyEffectorTypeLingTongShenFa
	// case types.LingTongDevSysTypeLingTi:
	// 	propertyEffectorType = playerpropertytypes.PlayerPropertyEffectorTypeLingTongXianTi
	// case types.LingTongDevSysTypeLingYi:
	// 	propertyEffectorType = playerpropertytypes.PlayerPropertyEffectorTypeLingTongWing
	// case types.LingTongDevSysTypeLingYu:
	// 	propertyEffectorType = playerpropertytypes.PlayerPropertyEffectorTypeLingTongLingYu
	// }
	// propertyManager.UpdateBattleProperty(propertyEffectorType.Mask())
	return
}

//变更灵童属性属性
func LingTongDevSelfPropertyChanged(pl player.Player, classType types.LingTongDevSysType) {
	//同步属性
	lingTongDataManager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	//写注册器会重复引用
	propertyEffectorType := classType.GetLingTongPropertyEffectType()

	lingTongDataManager.UpdateBattleProperty(propertyEffectorType.Mask())
	return
}

//灵童养成类进阶判断
func LingTongDevAdvanced(pl player.Player, curTimesNum int32, curBless int32, lingTongDevTemplate gametemplate.LingTongDevTemplate) (sucess bool, pro int32, randBless int32, addTimes int32, isDouble bool) {
	classType := lingTongDevTemplate.GetClassType()
	vipType := classType.GetVipType()
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, vipType, lingTongDevTemplate.GetTimesMin(), lingTongDevTemplate.GetTimesMax())
	updateRate := lingTongDevTemplate.GetUpdateWfb()
	blessMax := lingTongDevTemplate.GetZhuFuMax()
	addMin := lingTongDevTemplate.GetAddMin()
	addMax := lingTongDevTemplate.GetAddMax() + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	extralAddTimes := int32(0)
	welfareAdvancedType, ok := welfaretypes.LingTongDevTypeToAdvancedType(classType)
	if ok {
		isDouble, extralAddTimes = welfarelogic.IsCanAdvancedBlessCrit(welfareAdvancedType)
	}
	addTimes = int32(1)
	if isDouble {
		addTimes += extralAddTimes
		randBless += lingTongDevTemplate.GetAddMax()
	}

	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//灵童养成类皮肤升星判断
func LingTongDevSkinUpstar(curTimesNum int32, curBless int32, lingTongDevUpstarTemplate gametemplate.LingTongDevUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := lingTongDevUpstarTemplate.GetTimesMin()
	timesMax := lingTongDevUpstarTemplate.GetTimesMax()
	updateRate := lingTongDevUpstarTemplate.GetUpdateWfb()
	blessMax := lingTongDevUpstarTemplate.GetZhuFuMax()
	addMin := lingTongDevUpstarTemplate.GetAddMin()
	addMax := lingTongDevUpstarTemplate.GetAddMax() + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//灵童养成类通灵
func LingTongDevTongLing(curTimesNum int32, curBless int32, lingTongDevTongLingTemplate gametemplate.LingTongDevTongLingTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := lingTongDevTongLingTemplate.GetTimesMin()
	timesMax := lingTongDevTongLingTemplate.GetTimesMax()
	updateRate := lingTongDevTongLingTemplate.GetUpdateWfb()
	blessMax := lingTongDevTongLingTemplate.GetZhuFuMax()
	addMin := lingTongDevTongLingTemplate.GetAddMin()
	addMax := lingTongDevTongLingTemplate.GetAddMax() + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//灵童养成类喂养判断
func LingTongDevPeiYang(curTimesNum int32, curBless int32, lingTongDevPeiYangTemplate gametemplate.LingTongDevPeiYangTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := lingTongDevPeiYangTemplate.GetTimesMin()
	timesMax := lingTongDevPeiYangTemplate.GetTimesMax()
	updateRate := lingTongDevPeiYangTemplate.GetUpdateWfb()
	blessMax := lingTongDevPeiYangTemplate.GetZhuFuMax()
	addMin := lingTongDevPeiYangTemplate.GetAddMin()
	addMax := lingTongDevPeiYangTemplate.GetAddMax() + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//灵童养成类幻化丹培养判断
func LingTongDevHunaHuaFeed(curTimesNum int32, curBless int32, lingTongDevHuanHuaTemplate gametemplate.LingTongDevHuanHuaTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := lingTongDevHuanHuaTemplate.GetTimesMin()
	timesMax := lingTongDevHuanHuaTemplate.GetTimesMax()
	updateRate := lingTongDevHuanHuaTemplate.GetUpdateWfb()
	blessMax := lingTongDevHuanHuaTemplate.GetZhuFuMax()
	addMin := lingTongDevHuanHuaTemplate.GetAddMin()
	addMax := lingTongDevHuanHuaTemplate.GetAddMax() + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//灵童养成类祝福丹判断
func LingTongDevEatZhuFuDan(pl player.Player, curTimesNum, curBless, randBless int32, lingTongDevTemplate gametemplate.LingTongDevTemplate) (sucess bool, pro int32) {
	classType := lingTongDevTemplate.GetClassType()
	vipType := classType.GetVipType()
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, vipType, lingTongDevTemplate.GetTimesMin(), lingTongDevTemplate.GetTimesMax())

	updateRate := lingTongDevTemplate.GetUpdateWfb()
	blessMax := lingTongDevTemplate.GetZhuFuMax()
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}
