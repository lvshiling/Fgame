package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	tradelogic "fgame/fgame/game/trade/logic"
	"fgame/fgame/game/trade/trade"
)

//上传返还
func sellItem(target event.EventTarget, data event.EventData) (err error) {
	tradeItemObject, ok := target.(*trade.TradeItemObject)
	if !ok {
		return
	}
	playerId := tradeItemObject.GetPlayerId()
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)

	if p == nil {
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)
	playerTradeItem := message.NewScheduleMessage(onPlayerSellItem, ctx, tradeItemObject, nil)
	p.Post(playerTradeItem)

	return
}

func onPlayerSellItem(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	tradeItemObject := result.(*trade.TradeItemObject)
	tradelogic.OnPlayerSellItem(pl, tradeItemObject)
	return nil
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeSellItem, event.EventListenerFunc(sellItem))
}
