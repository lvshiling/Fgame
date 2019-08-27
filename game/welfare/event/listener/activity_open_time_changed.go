package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	"fgame/fgame/game/welfare/pbutil"
	"fgame/fgame/game/welfare/welfare"
)

//活动开服时间变化
func activityOpenTimeChanged(target event.EventTarget, data event.EventData) (err error) {
	alList := player.GetOnlinePlayerManager().GetAllPlayers()
	openTime := welfare.GetWelfareService().GetServerStartTime()
	mergeTime := welfare.GetWelfareService().GetServerMergeTime()
	for _, pl := range alList {
		scMsg := pbutil.BuildSCOpenActivityOpenServerTimeChanged(openTime, mergeTime)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeActivityOpenTimeChanged, event.EventListenerFunc(activityOpenTimeChanged))
}
