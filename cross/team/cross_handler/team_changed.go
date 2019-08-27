package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"
	teamtypes "fgame/fgame/game/team/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_TEAM_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerTeamChanged))
}

//玩家队伍变化
func handlePlayerTeamChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("team:玩家队伍变化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siPlayerTeamSync := msg.(*crosspb.SIPlayerTeamSync)
	teamId := siPlayerTeamSync.GetTeamData().GetTeamId()
	teamName := siPlayerTeamSync.GetTeamData().GetTeamName()
	teamPurpose := siPlayerTeamSync.GetTeamData().GetTeamPurpose()
	err = playerTeamChanged(tpl, teamId, teamName, teamPurpose)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("team:玩家队伍变化,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:玩家队伍变化,完成")
	return nil

}

//队伍变化
func playerTeamChanged(pl *player.Player, teamId int64, teamName string, teamPurpose int32) (err error) {
	pl.SyncTeam(teamId, teamName, teamtypes.TeamPurposeType(teamPurpose))
	return
}
