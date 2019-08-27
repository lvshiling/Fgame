package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
)

//九霄城战城复活占领停止
func allianceSceneReliveOccupyStop(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancescene.AllianceSceneData)
	s := sd.GetScene()
	allianceId := data.(int64)

	scAllianceSceneReliveOccupyStop := pbutil.BuildSCAllianceSceneReliveOccupyStop(allianceId)
	for _, p := range s.GetAllPlayers() {
		p.SendMsg(scAllianceSceneReliveOccupyStop)
	}
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneReliveOccupyStop, event.EventListenerFunc(allianceSceneReliveOccupyStop))
}
