package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLaBaClearLog, command.CommandHandlerFunc(handleLaBaClearLog))
}

func handleLaBaClearLog(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理拉霸日志重置")

	err = clearLaBaLog(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理拉霸日志重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理拉霸日志重置完成")
	return
}

func clearLaBaLog(p scene.Player) (err error) {
	welfare.GetWelfareService().GMClearLog()
	return
}
