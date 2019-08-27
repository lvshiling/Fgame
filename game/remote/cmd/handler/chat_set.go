package handler

import (
	"fgame/fgame/game/chat/chat"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_CHAT_SET_TYPE), cmd.CmdHandlerFunc(handleChatSet))
}

func handleChatSet(msg proto.Message) (err error) {
	log.Info("cmd:请求聊天设置")
	cmdChatSet := msg.(*cmdpb.CmdChatSet)
	worldVipLevel := cmdChatSet.GetWorldVipLevel()
	worldLevel := cmdChatSet.GetWorldLevel()
	allianceLevel := cmdChatSet.GetAllianceLevel()
	allianceVipLevel := cmdChatSet.GetAllianceVipLevel()
	privateLevel := cmdChatSet.GetPrivateLevel()
	privateVipLevel := cmdChatSet.GetPrivateVipLevel()
	teamLevel := cmdChatSet.GetTeamLevel()
	teamVipLevel := cmdChatSet.GetTeamVipLevel()

	err = chatSet(worldVipLevel, worldLevel, allianceVipLevel, allianceLevel, privateVipLevel, privateLevel, teamVipLevel, teamLevel)
	if err != nil {
		log.WithFields(
			log.Fields{
				"worldVipLevel":    worldVipLevel,
				"worldLevel":       worldLevel,
				"allianceVipLevel": allianceVipLevel,
				"allianceLevel":    allianceLevel,
				"privateVipLevel":  privateVipLevel,
				"privateLevel":     privateLevel,
				"err":              err,
			}).Error("cmd:请求聊天设置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"worldVipLevel":    worldVipLevel,
			"worldLevel":       worldLevel,
			"allianceVipLevel": allianceVipLevel,
			"allianceLevel":    allianceLevel,
			"privateVipLevel":  privateVipLevel,
			"privateLevel":     privateLevel,
		}).Info("cmd:请求聊天设置，成功")
	return
}

func chatSet(worldVipLevel, worldLevel, allianceVipLevel, allianceLevel, privateVipLevel, privateLevel, teamVipLevel, teamLevel int32) (err error) {
	chat.GetChatService().ChatSet(worldVipLevel, worldLevel, allianceVipLevel, allianceLevel, privateVipLevel, privateLevel, teamVipLevel, teamLevel)
	return
}
