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
	command.Register(gmcommandtypes.CommandTypeArenaStop, command.CommandHandlerFunc(handleArenaStop))
}

func handleArenaStop(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:3v3竞技场停止匹配")
	pl := p.(player.Player)
	err = arenaStop(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:3v3竞技场停止匹配,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:3v3竞技场停止匹配,完成")
	return
}

func arenaStop(pl player.Player) (err error) {
	arenalogic.ArenaStopMatchSend(pl)

	return

}
