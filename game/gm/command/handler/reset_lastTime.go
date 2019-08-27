package handler

import (
	"fgame/fgame/game/emperor/emperor"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/onearena/onearena"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLastTime, command.CommandHandlerFunc(handleResetLastTime))
}

//重置lastTime
func handleResetLastTime(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理重置lastTime")

	err = resetLastTime(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理重置lastTime,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理重置lastTime完成")
	return
}

//重置lastTime
func resetLastTime(pl scene.Player) (err error) {
	//重置帝王的lastTime
	emperor.GetEmperorService().GMResetLastTime()
	//重置鲲的lastTime
	onearena.GetOneArenaService().GmResetLastTime()
	return
}
