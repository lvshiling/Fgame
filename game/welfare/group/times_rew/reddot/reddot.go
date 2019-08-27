package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	grouprewtypes "fgame/fgame/game/welfare/group/times_rew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeTimesRew, reddot.HandlerFunc(handleRedDotTimesRew))
}

//累抽红点
func handleRedDotTimesRew(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*grouprewtypes.TimesRewInfo)
	timesTempList := welfaretemplate.GetWelfareTemplateService().GetTimesRewTemplateByGorup(groupId)
	for _, timesTemp := range timesTempList {
		if timesTemp.VipLevel > pl.GetVip() {
			continue
		}

		//是否领取
		if !info.IsCanReceiveRewards(timesTemp.DrawTimes, timesTemp.VipLevel) {
			continue
		}
		isNotice = true
		return
	}

	return
}
