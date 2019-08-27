package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	"fgame/fgame/game/welfare/pbutil"
)

func refreshXunHuan(target event.EventTarget, data event.EventData) (err error) {
	groupIdList, ok := data.([]int32)
	if !ok {
		return
	}

	// 跨天广播
	alPlList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range alPlList {
		scMsg := pbutil.BuildSCOpenActivityXunHuanInfo(groupIdList)
		pl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeRefreshXunHuanActivity, event.EventListenerFunc(refreshXunHuan))
}
