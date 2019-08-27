package handler

import (
	coredirty "fgame/fgame/core/dirty"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_ALLIANCE_NOTICE_TYPE), cmd.CmdHandlerFunc(handleAllianceNoticeSet))
}

func handleAllianceNoticeSet(msg proto.Message) (err error) {
	log.Info("cmd:请求仙盟公告设置")
	cmdChatSet := msg.(*cmdpb.CmdAllainceNotice)
	allianceId := cmdChatSet.GetAllianceId()
	noticeStr := cmdChatSet.GetNoticeStr()

	err = allianceNoticeSet(allianceId, noticeStr)
	if err != nil {
		log.WithFields(
			log.Fields{
				"allianceId": allianceId,
				"noticeStr":  noticeStr,
				"err":        err,
			}).Error("cmd:请求仙盟公告设置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"allianceId": allianceId,
			"noticeStr":  noticeStr,
		}).Info("cmd:请求仙盟公告设置，成功")
	return
}

func allianceNoticeSet(allianceId int64, content string) (err error) {
	content = strings.TrimSpace(content)
	lenOfContent := len([]rune(content))

	if lenOfContent < alliancetypes.MinAllianceNoticeLen && lenOfContent > alliancetypes.MaxAllianceNoticeLen {
		log.WithFields(
			log.Fields{
				"content": content,
			}).Warn("alliance:处理仙盟公告,内容长度不合法")
		err = cmd.ErrorCodeCommonArgumentInvalid
		return
	}

	flag := coredirty.GetDirtyService().IsLegal(content)
	if !flag {
		log.WithFields(
			log.Fields{
				"content": content,
			}).Warn("alliance:处理仙盟公告,内容含有脏字")
		err = cmd.ErrorCodeCommonArgumentDirty
		return
	}

	al, err := alliance.GetAllianceService().GMChangeAllianceNotice(allianceId, content)
	if err != nil {
		return
	}

	//推送其他成员
	scAllianceNoticeBroadcast := pbutil.BuildSCAllianceNoticeBroadcast(content)
	for _, mem := range al.GetMemberList() {
		memPlayer := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if memPlayer == nil {
			continue
		}
		memPlayer.SendMsg(scAllianceNoticeBroadcast)
	}
	return
}
