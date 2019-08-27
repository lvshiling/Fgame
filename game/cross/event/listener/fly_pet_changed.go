package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//宠物变化
func flyPetChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	// playerPetChanged := pbutil.BuildPlayerPetChanged(pl.GetPetId())
	// pl.SendCrossMsg(playerPetChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowFlyPetChanged, event.EventListenerFunc(flyPetChanged))
}
