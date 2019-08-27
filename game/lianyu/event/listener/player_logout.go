package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lianyulogic "fgame/fgame/game/lianyu/logic"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家下线
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	flag := pl.IsLianYuLineUp()
	if !flag {
		return
	}
	lianyulogic.LianYuCancleLineUpSend(pl)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
