package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/alliance/pbutil"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_INVITE_MERGE_FEEDBACK_TYPE), dispatch.HandlerFunc(handleAllianceInviteMergeFeedback))
}

//处理仙盟邀请合并反馈
func handleAllianceInviteMergeFeedback(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟邀请合并反馈")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSAllianceInviteMergeFeedback)
	agree := csMsg.GetAgree()
	inviteAllianceId := csMsg.GetInviteAllianceId()
	clearDepot := csMsg.GetClearAllianceDepot()
	clearHuFu := csMsg.GetClearAllianceHuFu()

	if agree {
		if !clearDepot || !clearHuFu {
			log.WithFields(
				log.Fields{
					"playerId":         pl.GetId(),
					"inviteAllianceId": inviteAllianceId,
					"clearDepot":       clearDepot,
					"clearHuFu":        clearHuFu,
					"error":            err,
				}).Warn("alliance:处理仙盟邀请合并反馈,参数错误")
			playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
			return
		}
	}

	err = allianceInviteMergeFeedback(tpl, agree, inviteAllianceId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
				"error":            err,
			}).Error("alliance:处理仙盟邀请合并反馈,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":         pl.GetId(),
			"inviteAllianceId": inviteAllianceId,
		}).Debug("alliance:处理仙盟邀请合并反馈,完成")
	return nil

}

//仙盟邀请合并反馈
func allianceInviteMergeFeedback(pl player.Player, agree bool, inviteAllianceId int64) (err error) {
	inviteAl := alliance.GetAllianceService().GetAlliance(inviteAllianceId)
	if inviteAl == nil {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
			}).Warn("found：处理仙盟邀请合并反馈,邀请的仙盟不存在")
		playerlogic.SendSystemMessage(pl, lang.AllianceNoExist)
		return
	}

	al := alliance.GetAllianceService().GetAlliance(pl.GetAllianceId())
	if al == nil {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
			}).Warn("found：处理仙盟邀请合并反馈,不在仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}

	if al.GetFuMengZhuId() != pl.GetId() && al.GetAllianceMengZhuId() != pl.GetId() {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
			}).Warn("found：处理仙盟邀请合并反馈,不是盟主或副盟主")
		return
	}

	defer alliance.GetAllianceService().ClearAllianceMergeInvite(inviteAllianceId)
	inviteMengzhu := player.GetOnlinePlayerManager().GetPlayerById(inviteAl.GetAllianceMengZhuId())
	inviteFuMengzhu := player.GetOnlinePlayerManager().GetPlayerById(inviteAl.GetFuMengZhuId())
	if !agree {

		if inviteMengzhu != nil {
			//TODO:xzk 修改为回协议
			playerlogic.SendSystemMessage(inviteMengzhu, lang.AllianceInviteMergeRefuse, pl.GetName())
		}
		if inviteFuMengzhu != nil {
			//TODO:xzk 修改为回协议
			playerlogic.SendSystemMessage(inviteFuMengzhu, lang.AllianceInviteMergeRefuse, pl.GetName())
		}
		return
	}

	mergeNum := al.NumOfMembers()
	if inviteAl.IfFull(mergeNum) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mergeNum": mergeNum,
				"curNum":   inviteAl.NumOfMembers(),
			}).Warn("found：处理仙盟邀请合并反馈,仙盟人数过多，无法合并反馈！")
		playerlogic.SendSystemMessage(pl, lang.AllianceInviteAllianceMergeMemberToMuch)
		if inviteMengzhu != nil {
			//TODO:xzk 修改为回协议
			playerlogic.SendSystemMessage(inviteMengzhu, lang.AllianceInviteAllianceMergeMemberToMuch)
		}
		if inviteFuMengzhu != nil {
			//TODO:xzk 修改为回协议
			playerlogic.SendSystemMessage(inviteFuMengzhu, lang.AllianceInviteAllianceMergeMemberToMuch)
		}
		return
	}

	if alliancelogic.IfOnAllianceActivity(pl.GetAllianceId()) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"allianceId": pl.GetAllianceId(),
			}).Warn("found：处理仙盟邀请合并,仙盟当前处于活动中，无法同意合并")
		playerlogic.SendSystemMessage(pl, lang.AllianceOnAllianceActivity)
		return
	}

	if alliancelogic.IfOnAllianceActivity(inviteAllianceId) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"allianceId": pl.GetAllianceId(),
			}).Warn("found：处理仙盟邀请合并,仙盟当前处于活动中，无法同意合并")
		playerlogic.SendSystemMessage(pl, lang.AllianceOnAllianceActivity)
		return
	}

	err = alliance.GetAllianceService().AllianceMerge(inviteAllianceId, al.GetAllianceId())
	if err != nil {
		return
	}

	// 系统频道
	inviteAlName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(inviteAl.GetAllianceName()))
	alName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(al.GetAllianceName()))
	systemContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceInviteAllianceMergeNotice), inviteAlName, alName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(systemContent))

	// 仙盟频道
	allianceContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceInviteAllianceMergeAllianceNotice), alName)
	chatlogic.SystemBroadcastAlliance(inviteAl, chattypes.MsgTypeText, []byte(allianceContent))

	scMsg := pbutil.BuildSCAllianceInviteMergeFeedback()
	pl.SendMsg(scMsg)
	return
}
