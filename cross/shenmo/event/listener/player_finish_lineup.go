package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/player/player"
	shenmologic "fgame/fgame/cross/shenmo/logic"
	"fgame/fgame/cross/shenmo/pbutil"
	gameevent "fgame/fgame/game/event"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
)

//玩家完成排队
func shenMoFinishLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	playerId := data.(int64)

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		isLianYuLineUpSuccess := pbutil.BuildISShenMoLineUpSuccess()
		pl.SendMsg(isLianYuLineUpSuccess)
	}

	shenmologic.BroadShenMoLineUpChanged(-1, lineList)
	return
}

func init() {
	gameevent.AddEventListener(shenmoeventtypes.EventTypeShenMoPlayerLineUpFinish, event.EventListenerFunc(shenMoFinishLineUp))
}
