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
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_UNIGNORE_PLAYER_CHAT_TYPE), cmd.CmdHandlerFunc(handleUnignorePlayerChat))
}

func handleUnignorePlayerChat(msg proto.Message) (err error) {
	log.Info("cmd:请求解除禁默")
	cmdUnignorePlayerChat := msg.(*cmdpb.CmdUnignorePlayerChat)
	forbidPlayerId := cmdUnignorePlayerChat.GetPlayerId()

	err = unignoreChatPlayer(forbidPlayerId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("cmd:请求解除禁默，错误")
		return
	}
	log.Info("cmd:请求解除禁默，成功")
	return
}

func unignoreChatPlayer(forbidPlayerId int64) (err error) {
	p := player.GetOnlinePlayerManager().GetPlayerById(forbidPlayerId)
	if p == nil {
		//加载离线玩家数据
		err = unignoreOfflinePlayerChat(forbidPlayerId)
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)

	msg := message.NewScheduleMessage(onUnignorePlayerChat, ctx, nil, nil)
	p.Post(msg)
	return
}

func unignoreOfflinePlayerChat(forbidPlayerId int64) (err error) {
	offlinePlayer, err := pp.CreateOfflinePlayer(forbidPlayerId)
	if err != nil {
		return err
	}
	if offlinePlayer == nil {
		err = cmd.ErrorCodeCommonPlayerNoExist
		return err
	}
	offlinePlayer.UnignoreChat()
	return nil
}

func onUnignorePlayerChat(ctx context.Context, result interface{}, err error) error {
	sp := scene.PlayerInContext(ctx)
	p, ok := sp.(player.Player)
	if !ok {
		return nil
	}

	p.UnignoreChat()
	return nil
}
