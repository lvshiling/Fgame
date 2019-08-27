package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/pbutil"
	playerteam "fgame/fgame/game/team/player"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_APPLY_ALL_TYPE), dispatch.HandlerFunc(handleTeamNearJoinAll))
}

//处理一键申请信息
func handleTeamNearJoinAll(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理一键申请消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = teamNearJoinAll(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理一键申请消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理一键申请消息完成")
	return nil

}

//一键申请信息的逻辑
func teamNearJoinAll(pl player.Player) (err error) {
	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	// teamId := mananger.GetTeamId()
	teamId := pl.GetTeamId()
	//当前处于队伍状态
	if teamId != 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("team:当前处于组队状态")
		playerlogic.SendSystemMessage(pl, lang.TeamPlayerInTeam)
		return
	}

	//操作过于频繁
	sucess := mananger.IfCanApplyAllTime()
	if !sucess {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("team:操作过于频繁")
		playerlogic.SendSystemMessage(pl, lang.CommonOperFrequent)
		return
	}

	err = team.GetTeamService().TeamApplyAll(pl)
	if err != nil {
		return
	}

	//更新时间
	mananger.TeamApplyAllTime()
	scTeamApplyAll := pbutil.BuildSCTeamApplyAll()
	pl.SendMsg(scTeamApplyAll)
	return
}
