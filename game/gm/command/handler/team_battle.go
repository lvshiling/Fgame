package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTeamBattle, command.CommandHandlerFunc(handleTeamBattle))

}

func handleTeamBattle(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"teamId": teamId,
			}).Warn("gm:处理组队副本开始战斗")
		return
	}

	purpose := pl.GetTeamPurpose()
	if !purpose.Vaild() {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"teamId":  teamId,
				"purpose": purpose,
			}).Warn("gm:处理组队副本开始战斗")
		return
	}

	err = team.GetTeamService().TeamCopyStartBattle(pl)
	if err != nil {
		return
	}
	return
}
