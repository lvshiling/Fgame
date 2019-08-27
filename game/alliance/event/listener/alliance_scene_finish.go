package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
)

//九霄城战结束
func allianceSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sceneData := target.(alliancescene.AllianceSceneData)
	s := sceneData.GetScene()

	allianceId := sceneData.GetCurrentDefendAllianceId()
	allianceName := sceneData.GetCurrentDefendAllianceName()
	alliancelogic.AllianceWin(allianceId)
	alliancelogic.DoorRewFinish(sceneData)

	scMsg := pbutil.BuildSCAllianceSceneFinish(allianceId, allianceName)
	for _, p := range s.GetAllPlayers() {
		p.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneFinish, event.EventListenerFunc(allianceSceneFinish))
}
