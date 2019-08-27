package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_INVITATION_TYPE), dispatch.HandlerFunc(handleAllianceInvitation))
}

//处理仙盟邀请
func handleAllianceInvitation(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟邀请")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceInvitation := msg.(*uipb.CSAllianceInvitation)
	invationId := csAllianceInvitation.GetInvitationId()

	err = allianceInvitation(tpl, invationId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"invationId": invationId,
				"error":      err,
			}).Error("alliance:处理仙盟邀请,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"invationId": invationId,
		}).Debug("alliance:处理仙盟邀请,完成")
	return nil

}

//仙盟邀请
func allianceInvitation(pl player.Player, invitationId int64) (err error) {
	member := alliance.GetAllianceService().GetAllianceMember(pl.GetId())
	//邀请权限
	if member.IsPositionMember() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("found：处理仙盟邀请,权限不足")
		playerlogic.SendSystemMessage(pl, lang.AllianceInvitationPositionNotEnough)
		return
	}

	beInvitedPlayer := player.GetOnlinePlayerManager().GetPlayerById(invitationId)
	if beInvitedPlayer != nil {
		ctx := scene.WithPlayer(context.Background(), beInvitedPlayer)
		msg := message.NewScheduleMessage(inviteJoinAlliance, ctx, pl.GetId(), nil)
		beInvitedPlayer.Post(msg)
	}

	scMsg := pbutil.BuildSCAllianceInvitation(invitationId)
	pl.SendMsg(scMsg)
	return
}

func inviteJoinAlliance(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)
	inviteId := result.(int64)
	invitePl := player.GetOnlinePlayerManager().GetPlayerById(inviteId)

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeAlliance) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("found：处理仙盟邀请,被邀请者功能未开启")

		if invitePl != nil {
			playerlogic.SendSystemMessage(invitePl, lang.CommonFuncNoOpen)
		}
		return nil
	}
	al, err := alliance.GetAllianceService().Invitation(inviteId, pl.GetId(), pl.GetName(), pl.GetRole(), pl.GetSex(), pl.GetLevel(), pl.GetForce())
	if err != nil {
		return err
	}
	inviteMem := alliance.GetAllianceService().GetAllianceMember(inviteId)
	allianceId := int32(al.GetAllianceId())
	allianceName := al.GetAllianceObject().GetName()
	scMsg := pbutil.BuildSCAllianceInvitationNotice(allianceId, allianceName, inviteMem.GetMemberId(), inviteMem.GetName())
	pl.SendMsg(scMsg)

	if invitePl != nil {
		playerlogic.SendSystemMessage(invitePl, lang.AllianceSendInvitation)
	}
	return nil
}
