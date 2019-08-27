package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	"fgame/fgame/game/godsiege/godsiege"
	godsiegescene "fgame/fgame/game/godsiege/scene"
)

//玩家退出炼狱
func godSiegePlayerExitScene(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(godsiegescene.GodSiegeSceneData)
	godType := sd.GetGodType()

	num := sd.GetScenePlayerNum()
	godsiege.GetGodSiegeService().RemoveFirstLineUpPlayer(godType, num)
	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegePlayerExit, event.EventListenerFunc(godSiegePlayerExitScene))
}
