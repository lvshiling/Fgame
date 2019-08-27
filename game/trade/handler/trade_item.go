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
	tradelogic "fgame/fgame/game/trade/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TRADE_ITEM_TYPE), dispatch.HandlerFunc(handleTradeItem))
}

//处理交易物品
func handleTradeItem(s session.Session, msg interface{}) (err error) {
	log.Debug("trade:处理交易物品")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csTradeItem := msg.(*uipb.CSTradeItem)
	tradeId := csTradeItem.GetTradeId()

	err = tradeItem(tpl, tradeId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("trade:处理交易物品,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("trade:处理交易物品,完成")
	return nil
}

//交易
func tradeItem(pl player.Player, tradeId int64) (err error) {
	if !center.GetCenterService().IsTradeOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("trade:交易行关闭中")
		playerlogic.SendSystemMessage(pl, lang.TradeServiceClose)
		return
	}

	if pl.GetPrivilege() != playertypes.PrivilegeTypeNone && !center.GetCenterService().IsNeiWanJiaoYiOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("trade:交易行,没有权限")
		playerlogic.SendSystemMessage(pl, lang.TradeServiceNoPrivilege)
		return
	}
	err = tradelogic.TradeItem(pl, tradeId)
	if err != nil {
		return
	}

	return
}
