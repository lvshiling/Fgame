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
	"fgame/fgame/game/trade/pbutil"
	"fgame/fgame/game/trade/trade"
)

//上传返还
func tradeItem(target event.EventTarget, data event.EventData) (err error) {
	tradeOrderObject, ok := target.(*trade.TradeOrderObject)
	if !ok {
		return
	}
	playerId := tradeOrderObject.GetPlayerId()
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		return
	}

	scTradeItem := pbutil.BuildSCTradeItem(tradeOrderObject.GetTradeId())
	p.SendMsg(scTradeItem)
	ctx := scene.WithPlayer(context.Background(), p)
	playerTradeItem := message.NewScheduleMessage(onPlayerTradeItem, ctx, tradeOrderObject, nil)
	p.Post(playerTradeItem)

	return
}

func onPlayerTradeItem(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	tradeOrderObject := result.(*trade.TradeOrderObject)
	tradelogic.OnPlayerTradeItem(pl, tradeOrderObject)
	return nil
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeItem, event.EventListenerFunc(tradeItem))
}
