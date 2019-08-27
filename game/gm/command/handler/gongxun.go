package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeGongXun, command.CommandHandlerFunc(handlePlayerGongXun))
}

func handlePlayerGongXun(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置功勋值")

	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	gongXunStr := c.Args[0]
	gongXun, err := strconv.ParseInt(gongXunStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"gongXun": gongXunStr,
				"error":   err,
			}).Warn("gm:处理设置功德,gongXun不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	if gongXun < 0 {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"gongXun": gongXunStr,
				"error":   err,
			}).Warn("gm:处理设置功德,gongXun不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	err = addGongXun(pl, int32(gongXun))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:设置功德,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:设置功德,完成")
	return
}

func addGongXun(p scene.Player, gongXun int32) (err error) {
	if gongXun < 0 {
		return
	}
	pl := p.(player.Player)
	pl.SetShenMoGongXunNum(gongXun)

	return
}
