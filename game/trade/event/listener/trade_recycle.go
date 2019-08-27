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
)

//交易回收
func tradeRecycle(target event.EventTarget, data event.EventData) (err error) {
	p, ok := data.(player.Player)
	if !ok {
		return
	}

	ctx := scene.WithPlayer(context.Background(), p)
	playerTradeRecycle := message.NewScheduleMessage(onPlayerTradeRecycle, ctx, nil, nil)
	p.Post(playerTradeRecycle)
	return
}

func onPlayerTradeRecycle(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	tradelogic.OnPlayerTradeRecycle(pl)
	return nil
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeRecycle, event.EventListenerFunc(tradeRecycle))
}
