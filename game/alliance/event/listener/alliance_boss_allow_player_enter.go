package listener

import (
	"fgame/fgame/core/event"
	alliancebossscene "fgame/fgame/game/alliance/boss_scene"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
)

//允许仙盟成员进入仙盟boss
func allowPlayerEnterAllianceBoss(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancebossscene.AllianceBossSceneData)
	p := data.(player.Player)
	if sd == nil {
		return
	}
	s := sd.GetScene()
	if s == nil {
		return
	}
	bornPos := s.MapTemplate().GetBornPos()
	scenelogic.PlayerEnterScene(p, s, bornPos)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllowPlayerEnterAllianceBoss, event.EventListenerFunc(allowPlayerEnterAllianceBoss))
}
