package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//采集 采集完成
func playerCollectFinishWith(target event.EventTarget, data event.EventData) (err error) {
	spl, ok := target.(scene.Player)
	if !ok {
		return
	}
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*collecteventtypes.CollectFinishWithEventData)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeChengZhan {
		return
	}

	collectNpc := eventData.GetCollectNpc()
	allianceSd, ok := s.SceneDelegate().(alliancescene.AllianceSceneData)
	if !ok {
		return
	}

	allianceId := pl.GetAllianceId()
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	name := al.GetAllianceName()
	huFu := al.GetAllianceObject().GetHuFu()
	if !allianceSd.YuXiNpcCollectFinish(collectNpc, allianceId, name, huFu) {
		return
	}

	scMsg := pbutil.BuildSCAllianceSceneOccupyFinish(allianceId, name, huFu)
	s.BroadcastMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinishWith, event.EventListenerFunc(playerCollectFinishWith))
}
