package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	playerfourgod "fgame/fgame/game/fourgod/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//玩家退出四神遗迹场景
func fourGodPlayerExitScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	manager.ExitFourGod()
	return
}

func init() {
	gameevent.AddEventListener(fourgodeventtypes.EventTypeFourGodPlayerExit, event.EventListenerFunc(fourGodPlayerExitScene))
}
