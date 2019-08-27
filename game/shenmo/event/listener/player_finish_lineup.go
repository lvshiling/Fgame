package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	shenmologic "fgame/fgame/game/shenmo/logic"
)

//玩家完成排队
func shenMoFinishLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	// playerId := data.(int64)

	// 本服排队未处理
	// pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	// if pl != nil {
	// 	isLianYuLineUpSuccess := pbutil.BuildISShenMoLineUpSuccess()
	// 	pl.SendMsg(isLianYuLineUpSuccess)
	// }

	shenmologic.BroadShenMoLineUpChanged(-1, lineList)
	return
}

func init() {
	gameevent.AddEventListener(shenmoeventtypes.EventTypeShenMoPlayerLineUpFinish, event.EventListenerFunc(shenMoFinishLineUp))
}
