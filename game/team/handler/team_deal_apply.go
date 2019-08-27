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
	teamtypes "fgame/fgame/game/team/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_NEAR_JOIN_RESULT_TYPE), dispatch.HandlerFunc(handleTeamNearJoinResult))
}

//处理申请加入决策信息
func handleTeamNearJoinResult(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理申请加入决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamNearJoinResult := msg.(*uipb.CSTeamNearJoinResult)
	result := csTeamNearJoinResult.GetResult()
	applyId := csTeamNearJoinResult.GetApplyId()

	err = teamNearJoinResult(tpl, teamtypes.TeamResultType(result), applyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"applyId":  applyId,
				"error":    err,
			}).Error("team:处理申请加入决策消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理申请加入决策消息完成")
	return nil

}

//申请加入决策信息的逻辑
func teamNearJoinResult(pl player.Player, result teamtypes.TeamResultType, applyId int64) (err error) {
	if !result.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
			"applyId":  applyId,
		}).Warn("team:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	teamId := pl.GetTeamId()
	if teamId == 0 {
		return
	}
	applyPlayer := player.GetOnlinePlayerManager().GetPlayerById(applyId)
	if applyPlayer == nil {
		playerlogic.SendSystemMessage(pl, lang.TeamPlayerOff)
		return
	}

	_, err = team.GetTeamService().CaptainApplyChoose(pl, applyPlayer, result)
	if err != nil {
		return
	}

	scTeamNearJoinResult := pbutil.BuildSCTeamNearJoinResult(applyId)
	pl.SendMsg(scTeamNearJoinResult)
	return
}
