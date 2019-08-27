package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/onearena/pbutil"
	onearenascene "fgame/fgame/game/onearena/scene"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeLingChiFighting {
		return
	}

	sd := s.SceneDelegate()
	sceneData, ok := sd.(onearenascene.OneArenaSceneData)
	if !ok {
		return
	}
	spl := sceneData.GetOneArenaRobot()
	if spl == nil {
		return
	}

	scOneArenaRobot := pbutil.BuildSCOneArenaRobot(spl)
	pl.SendMsg(scOneArenaRobot)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
