package handler

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/game/player"
	pp "fgame/fgame/game/player/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_PRIVILEGE_SET_TYPE), cmd.CmdHandlerFunc(handlePrivilegeSet))
}

func handlePrivilegeSet(msg proto.Message) (err error) {
	log.Debug("cmd:权限设置")
	cmdPrivilegeSet := msg.(*cmdpb.CmdPrivilegeSet)
	playerId := cmdPrivilegeSet.GetPlayerId()
	privilege := cmdPrivilegeSet.GetPrivilege()
	privilegeType := types.PrivilegeType(privilege)
	if !privilegeType.Valid() {
		err = cmd.ErrorCodeCommonPlayerNoExist
		log.WithFields(
			log.Fields{
				"playerId":  playerId,
				"privilege": privilege,
			}).Warn("cmd:请求权限设置,参数错误")
		return
	}
	err = privilegeSet(playerId, privilegeType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  playerId,
				"privilege": privilege,
				"err":       err,
			}).Error("cmd:请求权限设置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":  playerId,
			"privilege": privilege,
		}).Debug("cmd:请求权限设置,成功")
	return
}

func privilegeSet(playerId int64, privilege types.PrivilegeType) (err error) {
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		//加载离线玩家数据
		err = privilegeSetOfflinePlayer(playerId, privilege)
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)

	msg := message.NewScheduleMessage(onPrivilegeSet, ctx, privilege, nil)
	p.Post(msg)
	return
}

func onPrivilegeSet(ctx context.Context, result interface{}, err error) error {
	sp := scene.PlayerInContext(ctx)
	p, ok := sp.(player.Player)
	if !ok {
		return nil
	}
	privilege := result.(types.PrivilegeType)
	p.SetPrivilege(privilege)
	return nil
}

func privilegeSetOfflinePlayer(playerId int64, privilege types.PrivilegeType) (err error) {
	offlinePlayer, err := pp.CreateOfflinePlayer(playerId)
	if err != nil {
		return err
	}
	if offlinePlayer == nil {
		err = cmd.ErrorCodeCommonPlayerNoExist
		return err
	}
	offlinePlayer.SetPrivilege(privilege)
	return
}
