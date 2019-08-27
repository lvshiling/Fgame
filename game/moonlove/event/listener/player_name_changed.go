package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	moonlovelogic "fgame/fgame/game/moonlove/logic"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家名字变化
func playerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeYueXiaQingYuan {
		return
	}

	moonLoveDelegate := s.SceneDelegate()
	moonLoveDelegate.(moonlovelogic.MoonloveSceneData).PlayerNameChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerNameChanged, event.EventListenerFunc(playerNameChanged))
}
