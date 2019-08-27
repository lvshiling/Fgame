package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_TEAMCOPY_START_BATTLE_TYPE), dispatch.HandlerFunc(handleTeamCopyStartBattle))
}

//处理组队副本开始战斗
func handleTeamCopyStartBattle(s session.Session, msg interface{}) (err error) {
	log.Debug("teamcopy:处理组队副本开始战斗")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isTeamCopyStartBattle := msg.(*crosspb.ISTeamCopyStartBattle)
	sucess := isTeamCopyStartBattle.GetSucess()

	err = teamCopyStartBattle(tpl, sucess)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sucess":   sucess,
			}).Error("teamcopy:处理组队副本开始战斗,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"sucess":   sucess,
		}).Debug("teamcopy:处理组队副本开始战斗,完成")
	return nil

}

//开始战斗
func teamCopyStartBattle(pl player.Player, sucess bool) (err error) {
	if !sucess {
		err = team.GetTeamService().TeamCopyFailed(pl)
		return
	}
	err = team.GetTeamService().TeamCopySucess(pl)
	return
}
