package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeRobotClear, command.CommandHandlerFunc(handleRobotClear))
}

func handleRobotClear(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:清除机器人")

	//TODO 修改物品数量
	err = robotClear(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:清除机器人,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:清除机器人,完成")
	return
}

func robotClear(pl scene.Player) (err error) {
	robot.GetRobotService().GMClear()
	return
}
