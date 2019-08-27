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
	processor.Register(codec.MessageType(uipb.MessageType_CS_TRADE_INFO_LIST_TYPE), dispatch.HandlerFunc(handleTradeInfoList))
}

//处理交易列表
func handleTradeInfoList(s session.Session, msg interface{}) (err error) {
	log.Debug("trade:处理交易列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = tradeInfoList(tpl)
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
		}).Debug("trade:处理交易列表完成")
	return nil
}

var (
	maxItems = 300
)

//处理交易列表
func tradeInfoList(pl player.Player) (err error) {
	if !center.GetCenterService().IsTradeOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("trade:交易行关闭中")
		playerlogic.SendSystemMessage(pl, lang.TradeServiceClose)
		return
	}
	globalTradeList := trade.GetTradeService().GetGlobalTradeList()
	if len(globalTradeList) == 0 {
		scTradeInfoList := pbutil.BuildSCTradeInfoList(globalTradeList, 1, 1)
		pl.SendMsg(scTradeInfoList)
		return
	}
	totalPage := (len(globalTradeList)-1)/maxItems + 1
	for i := 0; i < totalPage; i++ {
		firstIndex := i * maxItems
		endIndex := (i + 1) * maxItems
		if endIndex > len(globalTradeList) {
			endIndex = len(globalTradeList)
		}
		scTradeInfoList := pbutil.BuildSCTradeInfoList(globalTradeList[firstIndex:endIndex], int32(i+1), int32(totalPage))
		pl.SendMsg(scTradeInfoList)
	}
	return
}
