package listener

import (
	"fgame/fgame/core/event"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
)

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyLearnUseItem, event.EventListenerFunc(playerBabyLearnUseItem))
}

func playerBabyLearnUseItem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*babyeventtypes.PlayerBabyLearnUseItemEventData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetItemMap())
}
