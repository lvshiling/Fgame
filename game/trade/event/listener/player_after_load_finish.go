package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	tradelogic "fgame/fgame/game/trade/logic"
	"fgame/fgame/game/trade/trade"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	//发货
	unfinishOrderList := trade.GetTradeService().GetUnfinishOrderList(p)
	for _, order := range unfinishOrderList {
		tradelogic.OnPlayerTradeItem(p, order)
	}
	sellList := trade.GetTradeService().GetSellList(p)
	for _, sellItem := range sellList {
		tradelogic.OnPlayerSellItem(p, sellItem)
	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
