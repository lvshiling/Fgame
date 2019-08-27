package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	gamecommon "fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_AGREE_INVITATION_TYPE), dispatch.HandlerFunc(handleAllianceAgreeInvitation))
}

//处理仙盟邀请
func handleAllianceAgreeInvitation(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟邀请信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceAgreeInvitation := msg.(*uipb.CSAllianceAgreeInvitation)
	memberId := csAllianceAgreeInvitation.GetMemberId()
	agree := csAllianceAgreeInvitation.GetAgree()

	err = allianceAgreeInvitation(tpl, memberId, agree)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"memberId": memberId,
				"agree":    agree,
				"error":    err,
			}).Error("alliance:处理仙盟邀请信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"memberId": memberId,
			"agree":    agree,
		}).Debug("alliance:处理仙盟邀请信息,完成")
	return nil

}

//处理仙盟邀请
func allianceAgreeInvitation(pl player.Player, memberId int64, agree bool) (err error) {
	beInvitationId := pl.GetId()
	beInvitationName := pl.GetName()
	beInvitationSex := pl.GetSex()
	beInvitationLevel := pl.GetLevel()
	beInvitationForce := pl.GetForce()
	beInvitationZhuanSheng := pl.GetZhuanSheng()
	beInvitationLingyuId := pl.GetLingyuInfo().AdvanceId
	beInvitationVip := pl.GetVip()

	al, err := alliance.GetAllianceService().AgreeAllianceInvitation(memberId, beInvitationId, agree)
	if err != nil {
		terr, ok := err.(gamecommon.Error)
		if !ok {
			return err
		}
		// 特殊处理
		if terr.Code() == lang.AllianceAlreadyFullApply {
			playerlogic.SendSystemMessage(pl, lang.AllianceAlreadyFullApply)
			agree = false
			err = nil
		} else {
			return err
		}
	}
	if agree {
		//同步用户数据
		alliance.GetAllianceService().SyncMemberInfo(beInvitationId, beInvitationName, beInvitationSex, beInvitationLevel, beInvitationForce, beInvitationZhuanSheng, beInvitationLingyuId, beInvitationVip)
		invitationIdMember := alliance.GetAllianceService().GetAllianceMember(beInvitationId)
		if invitationIdMember == nil {
			panic("alliance invitaion join: 成员应该存在")
		}
		scAllianceInfo := pbutil.BuildSCAllianceInfo(al, invitationIdMember)
		pl.SendMsg(scAllianceInfo)

		allianceId := al.GetAllianceId()
		allianceName := al.GetAllianceObject().GetName()
		scAllianceAgreeJoinToReply := pbutil.BuildSCAllianceAgreeJoinApplyToApply(allianceId, allianceName, agree)
		pl.SendMsg(scAllianceAgreeJoinToReply)
	}

	//通知邀请者
	memberPlayer := player.GetOnlinePlayerManager().GetPlayerById(memberId)
	if memberPlayer != nil {
		scAllianceAgreeInvitation := pbutil.BuildSCAllianceAgreeInvitationNotice(beInvitationId, beInvitationName, agree)
		memberPlayer.SendMsg(scAllianceAgreeInvitation)
	}

	scAllianceInvitation := pbutil.BuildSCAllianceAgreeInvitation(memberId, agree)
	pl.SendMsg(scAllianceInvitation)

	return
}
