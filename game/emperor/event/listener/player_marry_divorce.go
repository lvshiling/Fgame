package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/emperor/emperor"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
)

//玩家离婚
func playerMarryDivorce(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryDivorceEventData)
	if !ok {
		return
	}
	divorceType := eventData.GetDivorceType()
	if divorceType != marrytypes.MarryDivorceTypeForce {
		return
	}
	playerId := pl.GetId()
	spouseId := eventData.GetSpouseId()
	emperor.GetEmperorService().ResetEmperorSpouseName(playerId, spouseId)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryDivorce, event.EventListenerFunc(playerMarryDivorce))
}
