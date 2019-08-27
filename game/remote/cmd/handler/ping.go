package handler

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_PING_TYPE), cmd.CmdHandlerFunc(handlerPing))
}

func handlerPing(msg proto.Message) (err error) {
	log.Info("cmd:请求禁默")
	cmdPing := msg.(*cmdpb.CmdPing)
	platformId := cmdPing.GetPlatformId()
	serverId := cmdPing.GetServerId()
	if platformId != global.GetGame().GetPlatform() {
		return cmd.ErrorCodeCommonPlatformWrong
	}
	if serverId != global.GetGame().GetServerIndex() {
		return cmd.ErrorCodeCommonServerWrong
	}
	log.WithFields(
		log.Fields{
			"platformId": platformId,
			"serverId":   serverId,
		}).Info("cmd:请求ping，成功")
	return
}
