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

//交易市场上架商品成功
func tradeUpload(target event.EventTarget, data event.EventData) (err error) {
	refundTradeItemObj, ok := target.(*trade.TradeItemObject)
	if !ok {
		return
	}
	playerId := refundTradeItemObj.GetPlayerId()
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}

	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeTradUpload, 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeUpload, event.EventListenerFunc(tradeUpload))
}
