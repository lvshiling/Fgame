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
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_IGNORE_PLAYER_CHAT_TYPE), cmd.CmdHandlerFunc(handleIgnorePlayerChat))
}

func handleIgnorePlayerChat(msg proto.Message) (err error) {
	log.Info("cmd:请求禁默")
	cmdIgnorePlayerChat := msg.(*cmdpb.CmdIgnorePlayerChat)
	forbidPlayerId := cmdIgnorePlayerChat.GetPlayerId()
	forbidReason := cmdIgnorePlayerChat.GetForbidReason()
	forbidName := cmdIgnorePlayerChat.GetForbidName()
	forbidTime := cmdIgnorePlayerChat.GetForbidTime()
	err = ignoreChatPlayer(forbidPlayerId, forbidReason, forbidName, forbidTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("cmd:请求禁默，错误")
		return
	}
	log.Info("cmd:请求禁默，成功")
	return
}

func ignoreChatPlayer(forbidPlayerId int64, forbidReason string, forbidName string, forbidTime int64) (err error) {
	p := player.GetOnlinePlayerManager().GetPlayerById(forbidPlayerId)
	if p == nil {
		//加载离线玩家数据
		err = ignoreOfflinePlayerChat(forbidPlayerId, forbidReason, forbidName, forbidTime)
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)
	result := &ignorePlayerChatData{
		forbidReason: forbidReason,
		forbidName:   forbidName,
		forbidTime:   forbidTime,
	}
	msg := message.NewScheduleMessage(onIgnorePlayerChat, ctx, result, nil)
	p.Post(msg)
	return
}

func ignoreOfflinePlayerChat(forbidPlayerId int64, forbidReason string, forbidName string, forbidTime int64) (err error) {
	offlinePlayer, err := pp.CreateOfflinePlayer(forbidPlayerId)
	if err != nil {
		return err
	}
	if offlinePlayer == nil {
		err = cmd.ErrorCodeCommonPlayerNoExist
		return err
	}
	offlinePlayer.IgnoreChat(forbidReason, forbidName, forbidTime)
	return nil
}

type ignorePlayerChatData struct {
	forbidReason string
	forbidName   string
	forbidTime   int64
}

func onIgnorePlayerChat(ctx context.Context, result interface{}, err error) error {
	sp := scene.PlayerInContext(ctx)
	p, ok := sp.(player.Player)
	if !ok {
		return nil
	}
	data := result.(*ignorePlayerChatData)
	p.IgnoreChat(data.forbidReason, data.forbidName, data.forbidTime)
	return nil
}
