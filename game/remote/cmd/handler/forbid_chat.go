package handler

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/game/player"
	pp "fgame/fgame/game/player/player"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_FORBID_PLAYER_CHAT_TYPE), cmd.CmdHandlerFunc(handleForbidPlayerChat))
}

func handleForbidPlayerChat(msg proto.Message) (err error) {
	log.Info("cmd:请求禁言")
	cmdForbidPlayerChat := msg.(*cmdpb.CmdForbidPlayerChat)
	forbidPlayerId := cmdForbidPlayerChat.GetPlayerId()
	forbidReason := cmdForbidPlayerChat.GetForbidReason()
	forbidName := cmdForbidPlayerChat.GetForbidName()
	forbidTime := cmdForbidPlayerChat.GetForbidTime()
	err = forbidChatPlayer(forbidPlayerId, forbidReason, forbidName, forbidTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("cmd:请求禁言，错误")
		return
	}
	log.Info("cmd:请求禁言，成功")
	return
}

func forbidChatPlayer(forbidPlayerId int64, forbidReason string, forbidName string, forbidTime int64) (err error) {
	p := player.GetOnlinePlayerManager().GetPlayerById(forbidPlayerId)
	if p == nil {
		//加载离线玩家数据
		err = forbidOfflinePlayerChat(forbidPlayerId, forbidReason, forbidName, forbidTime)
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)
	result := &forbidPlayerChatData{
		forbidReason: forbidReason,
		forbidName:   forbidName,
		forbidTime:   forbidTime,
	}
	msg := message.NewScheduleMessage(onForbidPlayerChat, ctx, result, nil)
	p.Post(msg)
	return
}

func forbidOfflinePlayerChat(forbidPlayerId int64, forbidReason string, forbidName string, forbidTime int64) (err error) {
	offlinePlayer, err := pp.CreateOfflinePlayer(forbidPlayerId)
	if err != nil {
		return err
	}
	if offlinePlayer == nil {
		err = cmd.ErrorCodeCommonPlayerNoExist
		return err
	}
	offlinePlayer.ForbidChat(forbidReason, forbidName, forbidTime)
	return nil
}

type forbidPlayerChatData struct {
	forbidReason string
	forbidName   string
	forbidTime   int64
}

func onForbidPlayerChat(ctx context.Context, result interface{}, err error) error {
	sp := scene.PlayerInContext(ctx)
	p, ok := sp.(player.Player)
	if !ok {
		return nil
	}
	data := result.(*forbidPlayerChatData)
	p.ForbidChat(data.forbidReason, data.forbidName, data.forbidTime)
	return nil
}
