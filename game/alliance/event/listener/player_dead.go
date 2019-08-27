package listener

import (
	"fgame/fgame/core/event"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//旗子收集 死亡打断
func playerDead(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}

	chengWaiSceneData, ok := sd.(alliancescene.AllianceSceneData)
	if !ok {
		return
	}
	if chengWaiSceneData.GetCollectRelivePlayerId() == pl.GetId() {
		chengWaiSceneData.ClearReliveOccupy()
	}
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(playerDead))
}
