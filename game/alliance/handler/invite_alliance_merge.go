package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_INVITE_MERGE_TYPE), dispatch.HandlerFunc(handleAllianceInviteMerge))
}

//处理仙盟邀请合并
func handleAllianceInviteMerge(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟邀请合并")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSAllianceInviteMerge)
	inviteAllianceId := csMsg.GetInviteAllianceId()

	err = allianceInviteMerge(tpl, inviteAllianceId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
				"error":            err,
			}).Error("alliance:处理仙盟邀请合并,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":         pl.GetId(),
			"inviteAllianceId": inviteAllianceId,
		}).Debug("alliance:处理仙盟邀请合并,完成")
	return nil

}

//仙盟邀请合并
func allianceInviteMerge(pl player.Player, inviteAllianceId int64) (err error) {

	inviteAl := alliance.GetAllianceService().GetAlliance(inviteAllianceId)
	if inviteAl == nil {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
			}).Warn("alliance:处理仙盟邀请合并,邀请的仙盟不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//邀请权限
	member := alliance.GetAllianceService().GetAllianceMember(pl.GetId())
	if !member.IsMengZhu() && !member.IsFuMengZhu() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"Position": member.GetPosition(),
			}).Warn("alliance:处理仙盟邀请合并,权限不足")
		playerlogic.SendSystemMessage(pl, lang.AllianceInvitationPositionNotEnough)
		return
	}

	al := alliance.GetAllianceService().GetAlliance(pl.GetAllianceId())
	// if al.GetAllianceObject().GetCampType() != inviteAl.GetAllianceObject().GetCampType() {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":   pl.GetId(),
	// 			"camp":       al.GetAllianceObject().GetCampType(),
	// 			"inviteCamp": inviteAl.GetAllianceObject().GetCampType(),
	// 		}).Warn("found：处理仙盟邀请合并,仙盟阵营不同，无法合并！")
	// 	playerlogic.SendSystemMessage(pl, lang.AllianceCampNotSame)
	// 	return
	// }
	mergeNum := inviteAl.NumOfMembers()
	if al.IfFull(mergeNum) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mergeNum": mergeNum,
				"curNum":   al.NumOfMembers(),
			}).Warn("alliance:处理仙盟邀请合并,仙盟人数过多，无法合并！")
		playerlogic.SendSystemMessage(pl, lang.AllianceInviteAllianceMergeMemberToMuch)
		return
	}

	invitedMengzhu := player.GetOnlinePlayerManager().GetPlayerById(inviteAl.GetAllianceMengZhuId())

	invitedFuMengzhu := player.GetOnlinePlayerManager().GetPlayerById(inviteAl.GetFuMengZhuId())
	if invitedMengzhu == nil && invitedFuMengzhu == nil {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
			}).Warn("alliance:处理仙盟邀请合并,该仙盟盟主不在线，无法邀请")
		playerlogic.SendSystemMessage(pl, lang.AllianceInviteAllianceMergeMemberNotOnline)
		return
	}

	err = alliance.GetAllianceService().AllianceMergeInvite(pl.GetAllianceId(), inviteAllianceId)
	if err != nil {
		return
	}

	if alliancelogic.IfOnAllianceActivity(pl.GetAllianceId()) {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
			}).Warn("alliance:处理仙盟邀请合并,当前仙盟当前处于活动中，无法合并")
		playerlogic.SendSystemMessage(pl, lang.AllianceOnAllianceActivity)
		return
	}

	if alliancelogic.IfOnAllianceActivity(inviteAllianceId) {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"inviteAllianceId": inviteAllianceId,
			}).Warn("alliance:处理仙盟邀请合并,对方仙盟当前处于活动中，无法合并")
		playerlogic.SendSystemMessage(pl, lang.AllianceOnAllianceActivity)
		return
	}
	if invitedMengzhu != nil {
		inviteScMsg := pbutil.BuildSCAllianceInviteMergeNotice(al)
		invitedMengzhu.SendMsg(inviteScMsg)
	}
	if invitedFuMengzhu != nil {
		inviteScMsg := pbutil.BuildSCAllianceInviteMergeNotice(al)
		invitedFuMengzhu.SendMsg(inviteScMsg)
	}
	scMsg := pbutil.BuildSCAllianceInviteMerge(inviteAllianceId, inviteAl.GetAllianceName())
	pl.SendMsg(scMsg)
	return
}
