package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	"fgame/fgame/game/trade/trade"

	log "github.com/Sirupsen/logrus"
)

//交易退还
func tradeItemRefund(target event.EventTarget, data event.EventData) (err error) {
	tradeOrderObject, ok := target.(*trade.TradeOrderObject)
	if !ok {
		return
	}
	playerId := tradeOrderObject.GetPlayerId()
	//系统回收的
	if playerId == 0 {
		//TODO:zrc
		log.WithFields(
			log.Fields{
				"tradeId": tradeOrderObject.GetTradeId(),
			}).Info("trade:交易返还,系统回收的")
		return
	}
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	refundTitle := lang.GetLangService().ReadLang(lang.TradeItemRefundTitle)
	refundContent := lang.GetLangService().ReadLang(lang.TradeItemRefundContent)
	if p == nil {
		attachment := make(map[int32]int32)
		attachment[constanttypes.GoldItem] = int32(tradeOrderObject.GetGold())
		emaillogic.AddOfflineEmail(playerId, refundTitle, refundContent, attachment)
	} else {
		ctx := scene.WithPlayer(context.Background(), p)
		payerTradeItemRefund := message.NewScheduleMessage(onPlayerTradeItemRefund, ctx, tradeOrderObject, nil)
		p.Post(payerTradeItemRefund)
	}
	return
}

func onPlayerTradeItemRefund(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	tradeOrderObject := result.(*trade.TradeOrderObject)
	//TODO:zrc 判断是否足够物品
	refundTitle := lang.GetLangService().ReadLang(lang.TradeItemRefundTitle)
	refundContent := lang.GetLangService().ReadLang(lang.TradeItemRefundContent)
	attachment := make(map[int32]int32)
	attachment[constanttypes.GoldItem] = int32(tradeOrderObject.GetGold())
	emaillogic.AddEmail(pl, refundTitle, refundContent, attachment)
	return nil
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeItemRefund, event.EventListenerFunc(tradeItemRefund))
}
