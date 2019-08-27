package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
)

type reliveOccupyData struct {
	allianceId int64
	playerId   int64
}

//复活点正在占领
func allianceSceneReliveOccupying(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancescene.AllianceSceneData)
	eventData := data.(*allianceeventtypes.ReliveOccupyEventData)
	s := sd.GetScene()
	allianceId := eventData.GetAllianceId()
	playerId := eventData.GetPlayerId()

	scAllianceSceneOccupying := pbutil.BuildSCAllianceSceneReliveOccupying(allianceId, playerId)
	for _, p := range s.GetAllPlayers() {
		p.SendMsg(scAllianceSceneOccupying)
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneReliveOccupying, event.EventListenerFunc(allianceSceneReliveOccupying))
}
