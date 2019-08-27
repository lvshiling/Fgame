package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeCrossExit, command.CommandHandlerFunc(handleExitCross))
}

func handleExitCross(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:退出跨服")

	err = crossExit(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:跨服,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:跨服,完成")
	return
}

func crossExit(p scene.Player) (err error) {
	pl := p.(player.Player)
	crossSession := pl.GetCrossSession()
	if crossSession == nil {
		return
	}
	crossSession.Close(true)
	return

}
