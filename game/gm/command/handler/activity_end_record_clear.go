package handler

import (
	"fgame/fgame/game/activity/activity"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeActivityRecordClear, command.CommandHandlerFunc(handleEndRecordClear))
}

func handleEndRecordClear(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理日常活动记录重置")

	err = clearEndRecord(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理日常活动记录重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理日常活动记录重置完成")
	return
}

func clearEndRecord(p scene.Player) (err error) {
	activity.GetActivityService().GMClearEndRecord()
	return
}
