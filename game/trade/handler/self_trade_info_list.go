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
	"fgame/fgame/game/trade/pbutil"
	"fgame/fgame/game/trade/trade"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SELF_TRADE_INFO_LIST_TYPE), dispatch.HandlerFunc(handleSelfTradeInfoList))
}

//处理交易列表
func handleSelfTradeInfoList(s session.Session, msg interface{}) (err error) {
	log.Debug("trade:处理获取自己交易列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = selfTradeInfoList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("trade:处理获取自己交易列表,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("trade:处理获取自己交易列表完成")
	return nil
}

//处理交易列表
func selfTradeInfoList(pl player.Player) (err error) {
	if !center.GetCenterService().IsTradeOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("trade:交易行关闭中")
		playerlogic.SendSystemMessage(pl, lang.TradeServiceClose)
		return
	}
	tradeList := trade.GetTradeService().GetTradeList(pl)
	scSelfTradeInfoList := pbutil.BuildSCSelfTradeInfoList(tradeList)
	pl.SendMsg(scSelfTradeInfoList)
	return
}
