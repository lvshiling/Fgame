package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	pktypes "fgame/fgame/game/pk/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//玩家被杀
func playerKilled(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if !s.MapTemplate().IsWorld() {
		return
	}

	//陈楚楠说只有世界地图加红名值
	killedPlayer := data.(scene.Player)
	killedPlayerRedState := killedPlayer.GetPkRedState()

	switch killedPlayerRedState {
	case pktypes.PkRedStateInit:
		pl.Kill(true)
	default:
		pl.Kill(false)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypePlayerKilled, event.EventListenerFunc(playerKilled))
}
