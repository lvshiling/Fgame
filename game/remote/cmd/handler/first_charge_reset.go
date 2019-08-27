package handler

import (
	"fgame/fgame/game/charge/charge"
	"fgame/fgame/game/global"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	chargelogic "fgame/fgame/game/charge/logic"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_FIRST_CHARGE_RESET_TYPE), cmd.CmdHandlerFunc(handlerFirstChargeReset))
}

func handlerFirstChargeReset(msg proto.Message) (err error) {
	log.Info("cmd:请求首冲重置")

	err = firstChargeReset()
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("cmd:请求首冲重置")
		return
	}
	log.WithFields(
		log.Fields{}).Info("cmd:请求首冲重置")
	return
}

func firstChargeReset() (err error) {
	now := global.GetGame().GetTimeService().Now()
	flag := charge.GetChargeService().SetNewFirstChargeTime(now)
	if !flag {
		log.WithFields(
			log.Fields{}).Warn("cmd:首冲重置,重置失败")
		return cmd.ErrorCodeCommonFirstChargeFailed
	}
	chargelogic.BroadcastFirstCharge()
	return
}
