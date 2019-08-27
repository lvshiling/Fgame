package handler

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/game/player"
	pp "fgame/fgame/game/player/player"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/game/scene/scene"

	playerlogic "fgame/fgame/game/player/logic"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_FORBID_PLAYER_TYPE), cmd.CmdHandlerFunc(handleForbidPlayer))
}

func handleForbidPlayer(msg proto.Message) (err error) {
	log.Info("cmd:请求封号")
	cmdForbidPlayer := msg.(*cmdpb.CmdForbidPlayer)
	forbidPlayerId := cmdForbidPlayer.GetPlayerId()
	forbidReason := cmdForbidPlayer.GetForbidReason()
	forbidName := cmdForbidPlayer.GetForbidName()
	forbidTime := cmdForbidPlayer.GetForbidTime()
	err = forbidPlayer(forbidPlayerId, forbidReason, forbidName, forbidTime)
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

func forbidPlayer(forbidPlayerId int64, forbidReason string, forbidName string, forbidTime int64) (err error) {
	p := player.GetOnlinePlayerManager().GetPlayerById(forbidPlayerId)
	if p == nil {
		//加载离线玩家数据
		err = forbidOfflinePlayer(forbidPlayerId, forbidReason, forbidName, forbidTime)
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)
	result := &forbidPlayerData{
		forbidReason: forbidReason,
		forbidName:   forbidName,
		forbidTime:   forbidTime,
	}
	msg := message.NewScheduleMessage(onForbidPlayer, ctx, result, nil)
	p.Post(msg)

	//强制踢人
	player.GetOnlinePlayerManager().PlayerLeaveServer(p)
	//移除用户
	player.GetOnlinePlayerManager().RemovePlayer(p)

	return
}

func forbidOfflinePlayer(forbidPlayerId int64, forbidReason string, forbidName string, forbidTime int64) (err error) {
	offlinePlayer, err := pp.CreateOfflinePlayer(forbidPlayerId)
	if err != nil {
		return err
	}
	if offlinePlayer == nil {
		err = cmd.ErrorCodeCommonPlayerNoExist
		return err
	}
	offlinePlayer.Forbid(forbidReason, forbidName, forbidTime)
	return nil
}

type forbidPlayerData struct {
	forbidReason string
	forbidName   string
	forbidTime   int64
}

func onForbidPlayer(ctx context.Context, result interface{}, err error) error {
	sp := scene.PlayerInContext(ctx)
	p, ok := sp.(player.Player)
	if !ok {
		return nil
	}
	data := result.(*forbidPlayerData)
	p.Forbid(data.forbidReason, data.forbidName, data.forbidTime)
	//强制下线
	playerlogic.SendExceptionContentMessage(p, data.forbidReason)
	p.Close(nil)

	return nil
}
