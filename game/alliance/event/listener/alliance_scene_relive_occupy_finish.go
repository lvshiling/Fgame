package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//九霄城战城复活点占领完成
func allianceSceneReliveOccupyFinish(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancescene.AllianceSceneData)
	s := sd.GetScene()
	eventData := data.(*alliancescene.AllianceReliveCollectEventData)
	allianceId := eventData.GetAllianceId()
	playerId := eventData.GetPlayerId()
	playerName := ""
	//有两个地图 不能用场景去取
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p != nil {
		playerName = p.GetName()
	}
	allianceName := ""
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al != nil {
		allianceName = al.GetAllianceName()
	}
	scAllianceSceneReliveOccupyFinish := pbutil.BuildSCAllianceSceneReliveOccupyFinish(allianceId, allianceName, playerName)
	s.BroadcastMsg(scAllianceSceneReliveOccupyFinish)

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneReliveOccupyFinish, event.EventListenerFunc(allianceSceneReliveOccupyFinish))
}
