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
	command.Register(gmcommandtypes.CommandTypeBuffAdd, command.CommandHandlerFunc(handleBuffAdd))
}

func handleBuffAdd(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理设置buff")
	if len(c.Args) <= 0 {
		log.Warn("gm:处理设置buff,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	buffStr := c.Args[0]
	buffId, err := strconv.ParseInt(buffStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"buff":  buffStr,
			}).Warn("gm:处理设置buff,buff不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	//TODO 修改物品数量
	err = addBuff(pl, int32(buffId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"buff":  buffStr,
			}).Warn("gm:处理设置buff,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":   pl.GetId(),
			"buff": buffStr,
		}).Debug("gm:处理设置buff,完成")
	return
}

func addBuff(pl scene.Player, buff int32) (err error) {

	pl.AddBuff(buff, 0, 1, nil)
	return
}
