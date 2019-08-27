package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	relivelogic "fgame/fgame/game/relive/logic"
)

//加载完成后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)

	relivelogic.SyncReliveInfo(p)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
