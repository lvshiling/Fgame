package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeModel, command.CommandHandlerFunc(handleModel))
}

func handleModel(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置模型")

	modelStr := c.Args[0]
	model, err := strconv.ParseInt(modelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"model": modelStr,
			}).Warn("gm:设置模型,model不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	//TODO 修改物品数量
	err = setModel(pl, int32(model))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:设置模型,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:设置模型,完成")
	return
}

func setModel(pl scene.Player, model int32) (err error) {
	pl.SetModel(model)
	return
}
