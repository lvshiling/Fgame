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
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_UNFORBID_PLAYER_CHAT_TYPE), cmd.CmdHandlerFunc(handleUnforbidPlayerChat))
}

func handleUnforbidPlayerChat(msg proto.Message) (err error) {
	log.Info("cmd:请求解封禁言")
	cmdUnforbidPlayerChat := msg.(*cmdpb.CmdUnforbidPlayerChat)
	unforbidPlayerId := cmdUnforbidPlayerChat.GetPlayerId()

	err = unforbidPlayerChat(unforbidPlayerId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("cmd:请求解封禁言,错误")
		return
	}
	log.Info("cmd:请求解封禁言,成功")
	return
}

func unforbidPlayerChat(unforbidPlayerId int64) (err error) {
	p := player.GetOnlinePlayerManager().GetPlayerById(unforbidPlayerId)
	if p == nil {
		//加载离线玩家数据
		err = unforbidOfflinePlayerChat(unforbidPlayerId)
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)

	msg := message.NewScheduleMessage(onUnforbidPlayerChat, ctx, nil, nil)
	p.Post(msg)
	return
}

func unforbidOfflinePlayerChat(unforbidPlayerId int64) (err error) {
	offlinePlayer, err := pp.CreateOfflinePlayer(unforbidPlayerId)
	if err != nil {
		return err
	}
	if offlinePlayer == nil {
		err = cmd.ErrorCodeCommonPlayerNoExist
		return err
	}
	offlinePlayer.UnforbidChat()
	return nil
}

func onUnforbidPlayerChat(ctx context.Context, result interface{}, err error) error {
	sp := scene.PlayerInContext(ctx)
	p, ok := sp.(player.Player)
	if !ok {
		return nil
	}

	p.UnforbidChat()
	return nil
}
