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
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_NEAR_JOIN_TYPE), dispatch.HandlerFunc(handleTeamNearJoin))
}

//处理申请加入信息
func handleTeamNearJoin(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理申请加入消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamNearJoin := msg.(*uipb.CSTeamNearJoin)
	teamId := csTeamNearJoin.GetTeamId()

	err = teamNearJoin(tpl, teamId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"teamId":   teamId,
				"error":    err,
			}).Error("team:处理申请加入消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理申请加入消息完成")
	return nil

}

//申请加入信息的逻辑
func teamNearJoin(pl player.Player, teamId int64) (err error) {

	myTeamId := pl.GetTeamId()
	if myTeamId != 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"teamId":   myTeamId,
		}).Warn("team:当前处于组队状态")
		playerlogic.SendSystemMessage(pl, lang.TeamPlayerInTeam)
		return
	}

	resultCode := teamtypes.TeamApplyJoinCodeTypeSend
	flag, err := team.GetTeamService().JoinNearTeam(pl, teamId)
	//TODO 优化 之前走协议 后面又多了3v3 和组队副本
	switch err {
	case team.ErrorTeamInMatchJionFail,
		team.ErrorTeamInTeamCopyJionFail,
		team.ErrorTeamApplyJoinFuncNoOpen,
		team.ErrorTeamHouseIsBatting:
		return err
	case team.ErrorTeamPlayerFull:
		{
			resultCode = teamtypes.TeamApplyJoinCodeTypeFull
			break
		}
	case team.ErrorTeamDissolve:
		{
			resultCode = teamtypes.TeamApplyJoinCodeTypeDissolve
			break
		}
	case team.ErrorTeamJionByLeavedInCd:
		{
			return
		}
	}
	//自动审核
	if flag {
		resultCode = teamtypes.TeamApplyJoinCodeTypeSucess
	}

	scTeamNearJoin := pbutil.BuildSCTeamNearJoin(int32(resultCode), teamId)
	pl.SendMsg(scTeamNearJoin)
	return
}
