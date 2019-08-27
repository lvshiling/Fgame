package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/player/player"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
)

//玩家退出场景
func playerExitScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(*player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if pl.IsGuaJi() {
		_, flag := s.MapTemplate().GetMapType().GetGuaJiType()
		if !flag {
			return
		}
		pl.ExitGuaJi()
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(playerExitScene))
}
