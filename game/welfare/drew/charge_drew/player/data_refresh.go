package player

import (
	playercharge "fgame/fgame/game/charge/player"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	drewchargedrewtemplate "fgame/fgame/game/welfare/drew/charge_drew/template"
	drewchargedrewtypes "fgame/fgame/game/welfare/drew/charge_drew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeChargeDrew, playerwelfare.ActivityObjInfoRefreshHandlerFunc(chargeDrewRefreshInfo))
}

//充值抽奖-刷新
func chargeDrewRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime != 0 && now > endTime {
		return
	}

	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}
	groupInteface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(obj.GetGroupId())
	if groupInteface == nil {
		return
	}
	groupTemp := groupInteface.(*drewchargedrewtemplate.GroupTemplateChargeDrew)
	info := obj.GetActivityData().(*drewchargedrewtypes.LuckyChargeDrewInfo)
	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	chargeManager := obj.GetPlayer().GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)

	//跨天
	if !isSame {
		info.HadConvertTimes = 0
		convertRate := groupTemp.GetChargeDrewConvertRate()
		convertLimit := groupTemp.GetChargeDrewConvertLimit()
		minCycle := groupTemp.GetChargeDrewMinCycleTimes()
		info.CountLeftTimes(convertLimit, convertRate, minCycle)
		welfareManager.UpdateObj(obj)
	}

	isStartDay, _ := timeutils.IsSameDay(obj.GetStartTime(), now)
	//第一天同步数据
	if isStartDay {
		if info.GoldNum != int32(chargeManager.GetTodayChargeNum()) {
			addNum := int32(chargeManager.GetTodayChargeNum()) - info.GoldNum
			if addNum > 0 {
				info.GoldNum += addNum
				info.LeftConvertNum += addNum
				convertRate := groupTemp.GetChargeDrewConvertRate()
				convertLimit := groupTemp.GetChargeDrewConvertLimit()
				minCycle := groupTemp.GetChargeDrewMinCycleTimes()
				info.CountLeftTimes(convertLimit, convertRate, minCycle)
				welfareManager.UpdateObj(obj)
			}
		}
	}

	return
}
