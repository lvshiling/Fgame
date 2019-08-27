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
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_UNFORBID_PLAYER_TYPE), cmd.CmdHandlerFunc(handleUnforbidPlayer))
}

func handleUnforbidPlayer(msg proto.Message) (err error) {
	log.Info("cmd:请求解封")
	cmdUnforbidPlayer := msg.(*cmdpb.CmdUnforbidPlayer)
	unforbidPlayerId := cmdUnforbidPlayer.GetPlayerId()

	err = unforbidPlayer(unforbidPlayerId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("cmd:请求解封,错误")
		return
	}
	log.Info("cmd:请求解封,成功")
	return
}

func unforbidPlayer(unforbidPlayerId int64) (err error) {
	p := player.GetOnlinePlayerManager().GetPlayerById(unforbidPlayerId)
	if p == nil {
		//加载离线玩家数据
		err = unforbidOfflinePlayer(unforbidPlayerId)
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)

	msg := message.NewScheduleMessage(onUnforbidPlayer, ctx, nil, nil)
	p.Post(msg)
	return
}

func unforbidOfflinePlayer(unforbidPlayerId int64) (err error) {
	offlinePlayer, err := pp.CreateOfflinePlayer(unforbidPlayerId)
	if err != nil {
		return err
	}
	if offlinePlayer == nil {
		err = cmd.ErrorCodeCommonPlayerNoExist
		return err
	}
	offlinePlayer.Unforbid()
	return nil
}

func onUnforbidPlayer(ctx context.Context, result interface{}, err error) error {
	sp := scene.PlayerInContext(ctx)
	p, ok := sp.(player.Player)
	if !ok {
		return nil
	}

	p.Unforbid()
	return nil
}
