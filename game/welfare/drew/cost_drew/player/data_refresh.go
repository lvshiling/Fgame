package player

import (
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	drewcostdrewtemplate "fgame/fgame/game/welfare/drew/cost_drew/template"
	drewcostdrewtypes "fgame/fgame/game/welfare/drew/cost_drew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCostDrew, playerwelfare.ActivityObjInfoRefreshHandlerFunc(chargeDrewRefreshInfo))
}

//消费抽奖-刷新
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

	if !isSame {
		info := obj.GetActivityData().(*drewcostdrewtypes.LuckyCostDrewInfo)
		info.HadConvertTimes = 0

		groupInteface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(obj.GetGroupId())
		if groupInteface == nil {
			return
		}
		pl := obj.GetPlayer()
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		groupTemp := groupInteface.(*drewcostdrewtemplate.GroupTemplateCostDrew)
		convertRate := groupTemp.GetCostDrewConvertRate()
		convertLimit := groupTemp.GetCostDrewConvertLimit()
		minCycle := groupTemp.GetCostDrewMinCycleTimes()
		info.CountLeftTimes(convertLimit, convertRate, minCycle)
		welfareManager.UpdateObj(obj)
	}
	return
}
