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
	command.Register(gmcommandtypes.CommandTypeArenaNext, command.CommandHandlerFunc(handleArenaNext))
}

func handleArenaNext(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:3v3竞技场下一场匹配")
	pl := p.(player.Player)
	err = arenaNext(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:3v3竞技场下一场匹配,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:3v3竞技场下一场匹配,完成")
	return
}

func arenaNext(pl player.Player) (err error) {
	err = arenalogic.GMArenaNextSend(pl)
	if err != nil {
		return
	}
	return

}
