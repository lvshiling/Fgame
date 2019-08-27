package handler

import (
	feedbackfee "fgame/fgame/game/feedbackfee/feedbackfee"
	"fgame/fgame/game/player/dao"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_EXCHANGE_TYPE), cmd.CmdHandlerFunc(handleExchange))
}

func handleExchange(msg proto.Message) (err error) {
	log.Info("cmd:请求兑换")
	cmdExchange := msg.(*cmdpb.CmdExchange)
	exchangeId := cmdExchange.GetExchangeId()
	playerId := cmdExchange.GetPlayerId()
	code := cmdExchange.GetCode()
	money := cmdExchange.GetMoney()

	err = exchange(exchangeId, playerId, code, money)
	if err != nil {
		log.WithFields(
			log.Fields{
				"exchangeId": exchangeId,
				"playerId":   playerId,
				"code":       code,
				"money":      money,
				"err":        err,
			}).Error("cmd:请求兑换，错误")
		return
	}
	log.WithFields(
		log.Fields{
			"exchangeId": exchangeId,
			"playerId":   playerId,
			"code":       code,
			"money":      money,
		}).Info("cmd:请求兑换，成功")
	return
}

func exchange(exchangeId int64, playerId int64, code string, money int32) (err error) {
	pe, err := dao.GetPlayerDao().QueryById(playerId)
	if err != nil {
		return
	}
	if pe == nil {
		return cmd.ErrorCodeCommonPlayerNoExist
	}

	feedbackfee.GetFeedbackFeeService().CodeExchange(exchangeId, playerId, code, money)

	return
}
