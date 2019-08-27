package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
)

//九霄城战城门破了
func allianceSceneDoorBroke(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancescene.AllianceSceneData)
	s := sd.GetScene()
	currentDoor := sd.GetCurrentDoor()

	scAllianceSceneDoorBroke := pbutil.BuildSCAllianceSceneDoorBroke(currentDoor)
	for _, p := range s.GetAllPlayers() {
		p.SendMsg(scAllianceSceneDoorBroke)
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneDoorBroke, event.EventListenerFunc(allianceSceneDoorBroke))
}
