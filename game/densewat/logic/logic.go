package logic

import (
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

//协议兼容 神兽攻城
//金银密窟
func PlayerEnterDenseWat(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) (flag bool, err error) {
	return PlayerEnterDenseWatArgs(pl, activityTemplate, "")
}

func PlayerEnterDenseWatArgs(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	//进入跨服
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	timeTemp, _ := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
	if timeTemp == nil {
		return
	}
	endTime, _ := timeTemp.GetEndTime(now)
	denseEndTime := pl.GetDenseWatEndTime()
	if endTime != denseEndTime {
		pl.SetDenseWatEndTime(endTime)
	}
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeDenseWat)
	flag = true
	return
}
