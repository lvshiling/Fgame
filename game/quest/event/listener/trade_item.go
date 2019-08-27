package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	"fgame/fgame/game/trade/trade"
)

//交易市场购买
func tradeItem(target event.EventTarget, data event.EventData) (err error) {
	tradeOrderObject, ok := target.(*trade.TradeOrderObject)
	if !ok {
		return
	}
	playerId := tradeOrderObject.GetPlayerId()
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}

	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeTradItem, 0, tradeOrderObject.GetNum())
	return
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeItem, event.EventListenerFunc(tradeItem))
}
