package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/lineup/lineup"
	"fgame/fgame/cross/player/player"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
)

//退出场景
func playerExitScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(*player.Player)
	if !ok {
		return
	}

	s := pl.GetScene()
	crossType, ok := s.MapTemplate().GetMapType().ToCrossType()
	if !ok {
		return
	}

	lineup.GetLineupService().RemoveFirstLineUpPlayer(crossType, s.Id())
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(playerExitScene))
}
