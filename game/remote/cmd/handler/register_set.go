package handler

import (
	"fgame/fgame/game/register/register"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_REGISTER_SET_TYPE), cmd.CmdHandlerFunc(handleRegisterSet))
}

func handleRegisterSet(msg proto.Message) (err error) {
	log.Info("cmd:请求注册设置")
	cmdRegisterSet := msg.(*cmdpb.CmdRegisterSet)
	open := cmdRegisterSet.GetOpen()

	err = cmdRegister(open)
	if err != nil {
		log.WithFields(
			log.Fields{
				"open": open,

				"err": err,
			}).Error("cmd:请求注册设置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"open": open,
		}).Info("cmd:请求注册设置,成功")
	return
}

func cmdRegister(open int32) (err error) {
	if open != 0 {
		register.GetRegisterService().OpenRegister()
	} else {
		register.GetRegisterService().CloseRegister()
	}
	return
}
