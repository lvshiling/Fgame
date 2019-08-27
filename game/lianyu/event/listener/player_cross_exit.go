package listener

import (
	"fgame/fgame/core/event"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//玩家跨服断开
func playerCrossExit(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsLianYuLineUp() {
		return
	}
	pl.LianYuLineUp(false)
	//退出跨服
	// crosslogic.PlayerExitCross(pl)
	return
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossExit, event.EventListenerFunc(playerCrossExit))
}
