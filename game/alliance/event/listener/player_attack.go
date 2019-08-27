package listener

import (
	"fgame/fgame/core/event"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//旗子收集 攻击打断
func playerAttack(target event.EventTarget, data event.EventData) (err error) {
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

	asd, ok := sd.(alliancescene.AllianceSceneData)
	if !ok {
		return
	}
	if asd.GetCollectRelivePlayerId() == pl.GetId() {
		asd.ClearReliveOccupy()
	}
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectAttack, event.EventListenerFunc(playerAttack))
}
