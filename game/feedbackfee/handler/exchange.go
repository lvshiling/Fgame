package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	feedbackfeelogic "fgame/fgame/game/feedbackfee/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEEDBACK_FEE_EXCHANGE_TYPE), dispatch.HandlerFunc(handlerFeedbackExchange))
}

//处理兑换
func handlerFeedbackExchange(s session.Session, msg interface{}) (err error) {
	log.Debug("feedbackfee:处理获取兑换请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csFeedbackfeeExchange := msg.(*uipb.CSFeedbackFeeExchange)
	money := csFeedbackfeeExchange.GetExchangeMoney()
	if money <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"money":    money,
			}).Warn("feedbackfee:处理获取兑换请求,参数错误")
		return
	}
	err = exchange(tpl, money)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("feedbackfee:处理获取兑换请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("feedbackfee:处理获取兑换请求完成")

	return
}

//获取逆付费请求逻辑
func exchange(pl player.Player, money int32) (err error) {

	feedbackfeelogic.HandlePlayerExchange(pl, money)
	return
}
