package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"

	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	scScenePlayerData := pbutil.BuildSCScenePlayerData(p)
	p.SendMsg(scScenePlayerData)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
