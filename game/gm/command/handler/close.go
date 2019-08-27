package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeClose, command.CommandHandlerFunc(handleClose))
}

func handleClose(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理关闭")

	err = playerClose(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理婚车,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理关闭,完成")
	return
}

func playerClose(p scene.Player) (err error) {
	p.Close(nil)
	return
}
