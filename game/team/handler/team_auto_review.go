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
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_AUTO_REVIEW_TYPE), dispatch.HandlerFunc(handleTeamAutoReview))
}

//处理队伍自动审核
func handleTeamAutoReview(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理队伍自动审核")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamAutoReview := msg.(*uipb.CSTeamAutoReview)
	autoReview := csTeamAutoReview.GetAutoReview()

	err = teamAutoReview(tpl, autoReview)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"autoReview": autoReview,
				"error":      err,
			}).Error("team:处理队伍自动审核,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理队伍自动审核完成")
	return nil

}

//获取组队队伍自动审核的逻辑
func teamAutoReview(pl player.Player, autoReview bool) (err error) {
	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("team:队伍不存在")
		playerlogic.SendSystemMessage(pl, lang.TeamNoExist)
		return
	}

	err = team.GetTeamService().AutoReviewChoose(pl, autoReview)
	if err != nil {
		return
	}

	scTeamAutoReview := pbutil.BuildSCTeamAutoReview(autoReview)
	pl.SendMsg(scTeamAutoReview)
	return
}
