package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_INVITE_TYPE), dispatch.HandlerFunc(handleTeamInvite))
}

//处理邀请玩家信息
func handleTeamInvite(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理获取邀请玩家消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamInvite := msg.(*uipb.CSTeamInvite)
	invitedId := csTeamInvite.GetInvitedId()

	err = teamInvite(tpl, invitedId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"invitedId": invitedId,
				"error":     err,
			}).Error("team:处理获取邀请玩家消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理获取邀请玩家消息完成")
	return nil

}

//获取邀请玩家信息的逻辑
func teamInvite(pl player.Player, invitedId int64) (err error) {

	invitePlayer := player.GetOnlinePlayerManager().GetPlayerById(invitedId)
	//被邀请者已下线
	if invitePlayer == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"invitedId": invitedId,
			}).Warn("team:处理队伍邀请,玩家已经下线")
		playerlogic.SendSystemMessage(pl, lang.TeamPlayerOff)
		return
	}

	err = team.GetTeamService().InvitePlayer(pl, invitedId)
	if err != nil {
		return
	}

	scTeamInvite := pbutil.BuildSCTeamInvite(invitedId)
	pl.SendMsg(scTeamInvite)
	return
}
