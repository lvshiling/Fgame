package info

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/welfare/pbutil"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeFirstDrop, welfare.InfoGetHandlerFunc(handlerFirstKillInfo))
}

//boss首杀信息请求
func handlerFirstKillInfo(pl player.Player, groupId int32) (err error) {
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	record := welfare.GetWelfareService().GetBossFirstKillRecord(groupId)
	scMsg := pbutil.BuildSCOpenActivityGetInfo(groupId, startTime, endTime, record)
	pl.SendMsg(scMsg)
	return
}
