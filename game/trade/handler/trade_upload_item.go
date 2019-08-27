package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/center/center"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	tradelogic "fgame/fgame/game/trade/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TRADE_UPLOAD_ITEM_TYPE), dispatch.HandlerFunc(handleTradeUploadItem))
}

//处理交易上架
func handleTradeUploadItem(s session.Session, msg interface{}) (err error) {
	log.Debug("trade:处理交易上架")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csTradeUploadItem := msg.(*uipb.CSTradeUploadItem)
	bagTypeInt := csTradeUploadItem.GetBagType()
	index := csTradeUploadItem.GetIndex()
	num := csTradeUploadItem.GetNum()
	gold := csTradeUploadItem.GetGold()
	bagType := inventorytypes.BagType(bagTypeInt)

	if !bagType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("trade:处理交易上架,背包类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = tradeUploadItem(tpl, bagType, index, num, gold)
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
func tradeUploadItem(pl player.Player, typ inventorytypes.BagType, index int32, num int32, gold int32) (err error) {
	if !center.GetCenterService().IsTradeOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("trade:交易行关闭中")
		playerlogic.SendSystemMessage(pl, lang.TradeServiceClose)
		return
	}
	err = tradelogic.TradeUploadItem(pl, typ, index, num, gold)
	if err != nil {
		return
	}
	return
}
