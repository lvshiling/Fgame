package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
	zhenfaeventtypes "fgame/fgame/game/zhenfa/event/types"
)

func init() {
	gameevent.AddEventListener(zhenfaeventtypes.EventTypeZhenFaXianHuoShengJiUseItem, event.EventListenerFunc(playerZhenfaXianhuoShengjiUseItem))
}

func playerZhenfaXianhuoShengjiUseItem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*zhenfaeventtypes.PlayerZhenFaXianHuoShengJiUseItemEventData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetItemMap())
}
