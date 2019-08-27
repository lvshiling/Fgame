package handler

import (
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_KICK_OUT_PLAYER_TYPE), cmd.CmdHandlerFunc(handleKickOut))
}

func handleKickOut(msg proto.Message) (err error) {
	log.Info("cmd:请求踢人")
	kickOutPlayer := msg.(*cmdpb.CmdKickoutPlayer)
	playerId := kickOutPlayer.GetPlayerId()
	reason := kickOutPlayer.GetKickoutReason()
	err = kickoutPlayer(playerId, reason)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("cmd:请求替人，错误")
		return
	}
	log.Info("cmd:请求踢人，成功")
	return
}

func kickoutPlayer(playerId int64, reason string) (err error) {
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		return
	}
	//强制下线
	playerlogic.SendExceptionContentMessage(p, reason)
	p.Close(nil)
	player.GetOnlinePlayerManager().PlayerLeaveServer(p)
	//移除用户
	player.GetOnlinePlayerManager().RemovePlayer(p)

	return
}
