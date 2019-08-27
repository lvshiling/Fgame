package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	tradelogic "fgame/fgame/game/trade/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TRADE_WITHDRAW_ITEM_TYPE), dispatch.HandlerFunc(handleTradeWithDrawItem))
}

//处理交易下架
func handleTradeWithDrawItem(s session.Session, msg interface{}) (err error) {
	log.Debug("trade:处理交易下架")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csTradeWithDrawItem := msg.(*uipb.CSTradeWithDrawItem)
	tradeId := csTradeWithDrawItem.GetTradeId()

	err = tradeWithDrawItem(tpl, tradeId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("trade:处理交易上架,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("trade:处理交易上架完成")
	return nil
}

//交易上架
func tradeWithDrawItem(pl player.Player, tradeId int64) (err error) {
	if !center.GetCenterService().IsTradeOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("trade:交易行关闭中")
		playerlogic.SendSystemMessage(pl, lang.TradeServiceClose)
		return
	}
	err = tradelogic.TradeWithdraw(pl, tradeId)
	if err != nil {
		return
	}
	return
}
