package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_MEMBER_POS_TYPE), dispatch.HandlerFunc(handleTeamMemberPos))
}

//处理队友位置请求信息
func handleTeamMemberPos(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理队友位置请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = teamMemberPos(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理队友位置请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理队友位置请求消息完成")
	return nil

}

//队友位置请求信息的逻辑
func teamMemberPos(pl player.Player) (err error) {

	teamId := pl.GetTeamId()
	if teamId == 0 {
		return
	}
	teamData := team.GetTeamService().GetTeam(teamId)
	if teamData == nil {
		return
	}

	member, _ := teamData.GetMember(pl.GetId())
	if member == nil {
		return
	}

	scTeamMemberPos := pbutil.BuildSCTeamMemberPos(pl.GetId(), teamData)
	pl.SendMsg(scTeamMemberPos)
	return
}
