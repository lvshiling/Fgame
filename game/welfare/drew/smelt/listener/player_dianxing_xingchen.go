package listener

import (
	"fgame/fgame/core/event"
	dianxingeventtypes "fgame/fgame/game/dianxing/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
)

func init() {
	gameevent.AddEventListener(dianxingeventtypes.EventTypeDianXingUseItem, event.EventListenerFunc(playerDianxingUseItem))
}

func playerDianxingUseItem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*dianxingeventtypes.PlayerDianXingUseItemData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetItemMap())
}
