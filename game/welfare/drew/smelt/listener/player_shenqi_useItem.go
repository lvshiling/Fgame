package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shenqieventtypes "fgame/fgame/game/shenqi/event/types"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
)

func init() {
	gameevent.AddEventListener(shenqieventtypes.EventTypeShenQiUseItem, event.EventListenerFunc(playerShenqiUseItem))
}

func playerShenqiUseItem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*shenqieventtypes.PlayerShenQiUseItemEventData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetUseItemMap())
}
