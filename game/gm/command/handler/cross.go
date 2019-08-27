package handler

import (
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeCross, command.CommandHandlerFunc(handleCross))
}

func handleCross(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:跨服")

	err = cross(pl)
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

func cross(p scene.Player) (err error) {
	pl := p.(player.Player)
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeArena)
	return nil
}
