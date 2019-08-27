package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/emperor/emperor"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/player"
)

//玩家离婚决策
func playerMarryDivorceDeal(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryDivorceDealEventData)
	if !ok {
		return
	}
	agree := eventData.GetAgree()
	if !agree {
		return
	}

	spouseId := eventData.GetSpouseId()
	playerId := pl.GetId()
	emperor.GetEmperorService().ResetEmperorSpouseName(playerId, spouseId)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryDivorceDeal, event.EventListenerFunc(playerMarryDivorceDeal))
}
