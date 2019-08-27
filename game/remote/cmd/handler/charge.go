package handler

import (
	"fgame/fgame/game/charge/charge"
	chargetemplate "fgame/fgame/game/charge/template"
	"fgame/fgame/game/player/dao"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_CHARGE_TYPE), cmd.CmdHandlerFunc(handleCharge))
}

func handleCharge(msg proto.Message) (err error) {
	log.Info("cmd:请求充值")
	cmdCharge := msg.(*cmdpb.CmdCharge)
	orderId := cmdCharge.GetOrderId()
	playerId := cmdCharge.GetPlayerId()
	chargeId := cmdCharge.GetChargeId()
	money := cmdCharge.GetMoney()

	err = orderCharge(orderId, playerId, chargeId, money)
	if err != nil {
		log.WithFields(
			log.Fields{
				"orderId":  orderId,
				"playerId": playerId,
				"chargeId": chargeId,
				"money":    money,
				"err":      err,
			}).Error("cmd:请求充值，错误")
		return
	}
	log.WithFields(
		log.Fields{
			"orderId":  orderId,
			"playerId": playerId,
			"chargeId": chargeId,
			"money":    money,
		}).Info("cmd:请求充值，成功")
	return
}

func orderCharge(orderId string, playerId int64, chargeId int32, money int32) (err error) {
	pe, err := dao.GetPlayerDao().QueryById(playerId)
	if err != nil {
		return
	}
	if pe == nil {
		return cmd.ErrorCodeCommonPlayerNoExist
	}
	//验证其它参数
	chargeTemplate := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	if chargeTemplate == nil {
		err = cmd.ErrorCodeCommonArgumentInvalid
		return
	}
	flag, err := charge.GetChargeService().Charge(orderId, playerId, chargeId, money)
	if err != nil {
		return
	}
	if !flag {
		panic(fmt.Errorf("充值应该成功"))
	}

	return
}
