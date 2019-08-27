package listener

import (
	"fgame/fgame/core/event"
	playercharge "fgame/fgame/game/charge/player"
	gameevent "fgame/fgame/game/event"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	discounttaocantemplate "fgame/fgame/game/welfare/discount/taocan/template"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	hallonlinetypes "fgame/fgame/game/welfare/hall/online/types"
	investleveltypes "fgame/fgame/game/welfare/invest/level/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

// 玩家加载推送的信息用处

//加载完成后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	// 刷新
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	todayGoldNum := chargeManager.GetTodayChargeNum()
	todayCostNum := propertyManager.GetTodayCostNum()
	welfareManager.SyncFirstDayChargeRecord(int32(todayGoldNum))
	welfareManager.SyncFirstDayCostRecord(todayCostNum)

	timeList := welfaretemplate.GetWelfareTemplateService().GetAllActivityTimeTemplate()
	for _, timeTemp := range timeList {
		err = welfareManager.RefreshActivityData(timeTemp.GetOpenType(), timeTemp.GetOpenSubType())
		if err != nil {
			return
		}
	}

	//首充信息
	isFirst := welfareManager.IsFirstCharge()
	isReceive := welfareManager.IsReceiveFirstCharge()
	firstChagreNotice := pbutil.BuildSCOpenActivityFirstChargeNotice(isFirst, isReceive)
	pl.SendMsg(firstChagreNotice)

	//在线抽奖数据推送
	onlineTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeOnline)
	for _, tiemTemp := range onlineTimeTempList {
		groupId := tiemTemp.Group
		obj := welfareManager.GetOpenActivityIfNotCreate(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeOnline, groupId)
		info := obj.GetActivityData().(*hallonlinetypes.WelfareOnlineInfo)
		scMsg := pbutil.BuildSCOpenActivityWelfareOnlineDataNotice(groupId, info.DrawTimes)
		pl.SendMsg(scMsg)
	}

	//超值套餐信息
	//taoCanInfo(pl)    //超值套餐 旭东要求 屏掉处理

	//加载城战助威奖励
	loadAllianceCheerWin(pl)

	//循环活动
	loadXunHuanActivity(pl)

	//养鸡活动推送
	loadfeedbackChargeDevelopInfo(pl)

	//天劫塔冲刺活动推送
	loadTianJieTaInfo(pl)

	return
}

func taoCanInfo(pl player.Player) {
	isBuyHuiYaunPlus := false
	isBuyEquipGift := false
	isBuyInvestLevel := false

	// 会员
	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	isBuyHuiYaunPlus = huiyuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)

	//等级投资、装备许愿礼包
	investType := welfaretypes.OpenActivityTypeInvest
	investSubType := welfaretypes.OpenActivityInvestSubTypeLevel
	discountType := welfaretypes.OpenActivityTypeDiscount
	discountSubType := welfaretypes.OpenActivityDiscountSubTypeZhuanSheng
	taocanType := welfaretypes.OpenActivityTypeDiscount
	taocanSubType := welfaretypes.OpenActivityDiscountSubTypeTaoCan
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(taocanType, taocanSubType)
	for _, timeTemp := range timeTempList {
		if !welfarelogic.IsOnActivityTime(timeTemp.Group) {
			continue
		}

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(timeTemp.Group)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*discounttaocantemplate.GroupTemplateDiscountTaoCan)
		for _, relateGroupId := range timeTemp.GetRelationToGroupList() {
			relateTimeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(relateGroupId)
			relateObj := welfareManager.GetOpenActivity(relateGroupId)
			if relateObj == nil {
				continue
			}

			if relateTimeTemp.GetOpenType() == investType && relateTimeTemp.GetOpenSubType() == investSubType {
				investLevelType := investleveltypes.InvesetLevelTypeJunior
				info := relateObj.GetActivityData().(*investleveltypes.InvestLevelInfo)
				isBuyInvestLevel = info.IsBuy(investLevelType)
			}
			if relateTimeTemp.GetOpenType() == discountType && relateTimeTemp.GetOpenSubType() == discountSubType {
				giftIndex := groupTemp.GetEquipGiftIndex()
				info := relateObj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
				isBuyEquipGift = info.IsBuy(giftIndex)
			}
		}

		scMsg := pbutil.BuildSCOpenActivityTaoCanInfo(timeTemp.Group, isBuyHuiYaunPlus, isBuyEquipGift, isBuyInvestLevel)
		pl.SendMsg(scMsg)
	}

}

func loadAllianceCheerWin(pl player.Player) {
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		return
	}

	typ := welfaretypes.OpenActivityTypeAlliance
	subType := welfaretypes.OpenActivityAllianceSubTypeAlliance
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfare.GetWelfareService().IsAllianceWin(groupId, allianceId) {
			continue
		}

		//发邮件
		welfarelogic.AllianceCheerEndMail(pl, groupId)
	}

	return
}

func loadXunHuanActivity(pl player.Player) {
	groupIdList := welfarelogic.GetXunHuanGroupList()
	if len(groupIdList) == 0 {
		return
	}

	scMsg := pbutil.BuildSCOpenActivityXunHuanInfo(groupIdList)
	pl.SendMsg(scMsg)
}

// 养鸡推送
func loadfeedbackChargeDevelopInfo(pl player.Player) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
		if !checkFlag {
			return
		}

		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		err := welfareManager.RefreshActivityDataByGroupId(groupId)
		if err != nil {
			return
		}

		obj := welfareManager.GetOpenActivity(groupId)
		startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
		scMsg := pbutil.BuildSCOpenActivityFeedbackDevelopInfo(obj, groupId, startTime, endTime)
		pl.SendMsg(scMsg)
	}

}

// 天劫塔冲刺活动推送
func loadTianJieTaInfo(pl player.Player) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeRealm

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
		if !checkFlag {
			return
		}

		startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		realmObj := welfareManager.GetOpenActivity(groupId)
		timesList := welfare.GetWelfareService().GetReceiveTimesList(groupId)

		scOpenActivityRealmInfo := pbutil.BuildSCOpenActivityRealmInfo(realmObj, groupId, timesList, startTime, endTime)
		pl.SendMsg(scOpenActivityRealmInfo)
	}
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
