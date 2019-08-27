package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	coredirty "fgame/fgame/core/dirty"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_NOTICE_CHANGE_TYPE), dispatch.HandlerFunc(handleAllianceNotice))
}

//仙盟信息
func handleAllianceNotice(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理修改仙盟公告")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceNoticeChange := msg.(*uipb.CSAllianceNoticeChange)
	content := csAllianceNoticeChange.GetContent()

	err = allianceNotice(tpl, content)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"content":  content,
				"error":    err,
			}).Error("alliance:处理修改仙盟公告,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"content":  content,
		}).Debug("alliance:处理修改仙盟公告,完成")
	return nil

}

func allianceNotice(pl player.Player, content string) (err error) {
	content = strings.TrimSpace(content)
	lenOfContent := len([]rune(content))

	if lenOfContent < alliancetypes.MinAllianceNoticeLen && lenOfContent > alliancetypes.MaxAllianceNoticeLen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"content":  content,
			}).Warn("alliance:处理仙盟公告,内容长度不合法")

		playerlogic.SendSystemMessage(pl, lang.AllianceNoticeIllegal)
		return
	}

	flag := coredirty.GetDirtyService().IsLegal(content)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"content":  content,
			}).Warn("alliance:处理仙盟公告,内容含有脏字")

		playerlogic.SendSystemMessage(pl, lang.AllianceNoticeDirty)
		return
	}

	al, err := alliance.GetAllianceService().ChangeAllianceNotice(pl.GetId(), content)
	if err != nil {
		return
	}

	//推送其他成员
	scAllianceNoticeBroadcast := pbutil.BuildSCAllianceNoticeBroadcast(content)
	for _, mem := range al.GetMemberList() {
		if mem.GetMemberId() == pl.GetId() {
			continue
		}

		memPlayer := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if memPlayer == nil {
			continue
		}
		memPlayer.SendMsg(scAllianceNoticeBroadcast)
	}

	scAllianceNoticeChange := pbutil.BuildSCAllianceNoticeChange(content)
	pl.SendMsg(scAllianceNoticeChange)
	return
}
