package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//玩家进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()

	if pl.IsGuaJi() {
		guaJiType, flag := s.MapTemplate().GetMapType().GetGuaJiType()
		if !flag {
			return
		}
		pl.EnterGuaJi(guaJiType)
	}

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
