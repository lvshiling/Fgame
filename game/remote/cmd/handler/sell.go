package handler

import (
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/game/trade/trade"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_TRADE_SELL_TYPE), cmd.CmdHandlerFunc(handleTradeSell))
}

func handleTradeSell(msg proto.Message) (err error) {
	log.Info("cmd:请求交易通知")
	cmdTradeSell := msg.(*cmdpb.CmdTradeSell)
	playerId := cmdTradeSell.GetPlayerId()
	platform := cmdTradeSell.GetPlatform()
	serverId := cmdTradeSell.GetServerId()
	tradeId := cmdTradeSell.GetTradeId()
	gold := cmdTradeSell.GetGold()
	globalTradeId := cmdTradeSell.GetGlobalTradeId()
	buyPlayerId := cmdTradeSell.GetBuyPlayerId()
	buyPlatform := cmdTradeSell.GetBuyPlatform()
	buyServerId := cmdTradeSell.GetBuyServerId()
	buyPlayerName := cmdTradeSell.GetBuyPlayerName()

	err = tradeSell(playerId, tradeId, buyPlatform, buyServerId, buyPlayerId, buyPlayerName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":      playerId,
				"platform":      platform,
				"serverId":      serverId,
				"gold":          gold,
				"globalTradeId": globalTradeId,
				"buyPlayerId":   buyPlayerId,
				"buyPlatform":   buyPlatform,
				"buyPlayerName": buyPlayerName,
				"err":           err,
			}).Error("cmd:请求交易通知")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":      playerId,
			"platform":      platform,
			"serverId":      serverId,
			"gold":          gold,
			"globalTradeId": globalTradeId,
			"buyPlayerId":   buyPlayerId,
			"buyPlatform":   buyPlatform,
			"buyPlayerName": buyPlayerName,
		}).Info("cmd:请求交易通知，成功")
	return
}

func tradeSell(playerId int64, tradeId int64, buyPlatform int32, buyServerId int32, buyPlayerId int64, buyPlayerName string) (err error) {
	trade.GetTradeService().SellItem(playerId, tradeId, buyPlatform, buyServerId, buyPlayerId, buyPlayerName)

	return
}
