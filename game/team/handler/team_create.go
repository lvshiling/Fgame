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
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_CREATE_BY_PLAYER_TYPE), dispatch.HandlerFunc(handleTeamCreateByPlayer))
}

//处理玩家点击创建队伍信息
func handleTeamCreateByPlayer(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理获取玩家点击创建队伍消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = teamCreateByPlayer(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理获取玩家点击创建队伍消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理获取玩家点击创建队伍消息完成")
	return nil

}

//获取玩家点击创建队伍信息的逻辑
func teamCreateByPlayer(pl player.Player) (err error) {
	//玩家已经在队伍中了
	if pl.GetTeamId() != 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("team:当前已有队伍,不能创建队伍")
		playerlogic.SendSystemMessage(pl, lang.TeamPlayerInTeam)
		return
	}

	//创建队伍
	teamData, err := team.GetTeamService().CreateTeamByPlayer(pl, teamtypes.TeamPurposeTypeNormal)
	if err != nil {
		return
	}

	teamId := teamData.GetTeamId()
	teamName := teamData.GetTeamName()

	// mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	pl.SyncTeam(teamId, teamName, teamtypes.TeamPurposeTypeNormal)
	//推送队员信息
	scTeamGet := pbutil.BuildSCTeamGet(teamData, false, pl.GetId())
	pl.SendMsg(scTeamGet)
	return
}
