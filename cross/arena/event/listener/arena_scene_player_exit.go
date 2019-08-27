package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/arena"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//竞技场景退出
func arenaScenePlayerExit(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	active := data.(bool)
	if !active {
		//成员
		arena.GetArenaService().ArenaMemberOffline(pl)
	}
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaScenePlayerExit, event.EventListenerFunc(arenaScenePlayerExit))
}
