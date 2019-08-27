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
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_BY_LEAVED_TYPE), dispatch.HandlerFunc(handleTeamByLeaved))
}

//处理队长请离信息
func handleTeamByLeaved(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理队长请离消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamByLeaved := msg.(*uipb.CSTeamByLeaved)
	leavedId := csTeamByLeaved.GetLeavedId()

	err = teamByLeaved(tpl, leavedId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"leavedId": leavedId,
				"error":    err,
			}).Error("team:处理队长请离消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理队长请离消息完成")
	return nil
}

//队长请离信息的逻辑
func teamByLeaved(pl player.Player, leavedId int64) (err error) {
	playerId := pl.GetId()
	if playerId == leavedId {
		log.WithFields(log.Fields{
			"playerId": playerId,
			"leavedId": leavedId,
		}).Warn("team:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(log.Fields{
			"playerId": playerId,
			"leavedId": leavedId,
		}).Warn("team:队伍不存在")
		playerlogic.SendSystemMessage(pl, lang.TeamNoExist)
		return
	}

	_, _, err = team.GetTeamService().BeLeavedTeam(pl, leavedId)
	if err != nil {
		return
	}
	return
}
