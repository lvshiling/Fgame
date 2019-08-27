package listener

import (
	"fgame/fgame/core/event"
	lianyulogic "fgame/fgame/cross/lianyu/logic"
	"fgame/fgame/cross/lianyu/pbutil"
	"fgame/fgame/cross/player/player"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
)

//玩家完成排队
func lianYuFinishLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	playerId := data.(int64)

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		isLianYuLineUpSuccess := pbutil.BuildISLianYuLineUpSuccess()
		pl.SendMsg(isLianYuLineUpSuccess)
	}

	lianyulogic.BroadLianYuLineUpChanged(-1, lineList)
	return
}

func init() {
	gameevent.AddEventListener(lianyueventtypes.EventTypeLianYuPlayerLineUpFinish, event.EventListenerFunc(lianYuFinishLineUp))
}
