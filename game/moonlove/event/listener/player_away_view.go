package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	moonlovelogic "fgame/fgame/game/moonlove/logic"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//月下情缘移动，解除双人赏月
func playerAway(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	switch pl := bo.(type) {
	case player.Player:

		//月下情缘场景
		if pl.GetScene().MapTemplate().GetMapType() != scenetypes.SceneTypeYueXiaQingYuan {
			return
		}

		sceneData := pl.GetScene().SceneDelegate().(moonlovelogic.MoonloveSceneData)
		isCouple := sceneData.IsCouple(pl.GetId())
		if isCouple {
			sceneData.ReleaseCouple(pl.GetId())
		}
		break
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectMove, event.EventListenerFunc(playerAway))
}
