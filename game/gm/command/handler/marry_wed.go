package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMarryWed, command.CommandHandlerFunc(handleClearMarryWed))
}

func handleClearMarryWed(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空婚期")

	err = clearMarryWed(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空婚期,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空婚期,完成")
	return
}

func clearMarryWed(p scene.Player) (err error) {

	marry.GetMarryService().GmClearMarryWed()
	return
}
