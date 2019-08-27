package handler

import (
	arenalogic "fgame/fgame/game/arena/logic"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeArena, command.CommandHandlerFunc(handleArena))
}

func handleArena(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:3v3竞技场")
	pl := p.(player.Player)
	err = arena(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:3v3竞技场,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:3v3竞技场,完成")
	return
}

func arena(pl player.Player) (err error) {
	_, err = arenalogic.ArenaMatch(pl)
	if err != nil {
		return
	}
	return
}
