package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeCollectPoker, reddot.HandlerFunc(handleRedDotCollectRew))
}

//卡牌收集红点
func handleRedDotCollectRew(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	// if !welfarelogic.IsOnActivityTime(groupId) {
	// 	return
	// }

	// welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	// obj := welfareManager.GetOpenActivity(groupId)
	// if obj == nil {
	// 	return
	// }

	// info := obj.GetActivityData().(*timesrewtypes.TimesRewInfo)
	// timesTempList := welfaretemplate.GetWelfareTemplateService().GetTimesRewTemplateByGorup(groupId)
	// for _, timesTemp := range timesTempList {
	// 	if timesTemp.VipLevel > pl.GetVip() {
	// 		continue
	// 	}

	// 	//是否领取
	// 	if !info.IsCanReceiveRewards(timesTemp.DrawTimes, timesTemp.VipLevel) {
	// 		continue
	// 	}
	// 	isNotice = true
	// 	return
	// }

	return
}
