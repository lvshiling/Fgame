package handler

import (
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_MARRY_SET_TYPE), cmd.CmdHandlerFunc(handleMarrySet))
}

func handleMarrySet(msg proto.Message) (err error) {
	log.Info("cmd:请求结婚版本设置")
	// cmdMsg := msg.(*cmdpb.CmdMarrySet)
	// houtai := cmdMsg.GetHouTaiType()
	// houtaiType := marrytypes.MarryHoutaiType(houtai)
	// if !houtaiType.Valid() {
	// 	log.WithFields(log.Fields{
	// 		"houtaiType": houtai,
	// 	}).Error("cmd:请求结婚版本设置 类型设置错误")
	// 	return fmt.Errorf("cmd:请求结婚版本设置 类型设置错误")
	// }
	// marry.GetMarrySetService().ResetHouTaiType(houtaiType)
	// log.Info("cmd:请求结婚版本设置，成功")
	// scMarryBanquetSet := pbuitl.BuildSCMarryBanquetSetChangeMsg(houtaiType)
	// player.GetOnlinePlayerManager().BroadcastMsg(scMarryBanquetSet)
	return
}
