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
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/trade/pbutil"
	playertrade "fgame/fgame/game/trade/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TRADE_LOG_LIST_TYPE), dispatch.HandlerFunc(handleTradeLogList))
}

//处理交易日志
func handleTradeLogList(s session.Session, msg interface{}) (err error) {
	log.Debug("trade:处理交易日志列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = tradeLogList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("trade:处理交易日志,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("trade:处理交易日志列表完成")
	return nil
}

//处理交易日志
func tradeLogList(pl player.Player) (err error) {
	if !center.GetCenterService().IsTradeOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("trade:交易行关闭中")
		playerlogic.SendSystemMessage(pl, lang.TradeServiceClose)
		return
	}
	tradeManager := pl.GetPlayerDataManager(playertypes.PlayerTradeDataManagerType).(*playertrade.PlayerTradeManager)
	tradeLogList := tradeManager.GetTradeLogList()
	scTradeLogList := pbutil.BuildSCTradeLogList(tradeLogList)
	pl.SendMsg(scTradeLogList)
	return
}
