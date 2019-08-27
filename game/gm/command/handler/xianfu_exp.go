package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	xianfulogic "fgame/fgame/game/xianfu/logic"
	xianfutypes "fgame/fgame/game/xianfu/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXianfuExp, command.CommandHandlerFunc(handleXianfuExp))
}

func handleXianfuExp(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	log.Debug("gm:处理进入秘境仙府经验副本")
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": level,
				"error": err,
			}).Warn("gm:处理跳转秘境仙府经验副本,秘境仙府副本id不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	err = enterXianfuExp(pl, int32(level))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": level,
				"error": err,
			}).Warn("gm:处理跳转秘境仙府经验副本,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":    pl.GetId(),
			"level": level,
		}).Debug("gm:处理跳转秘境仙府经验副本,完成")
	return
}

func enterXianfuExp(pl player.Player, level int32) (err error) {
	xianfulogic.PlayerEnterXianFu(pl, xianfutypes.XianfuTypeExp, level)
	return
}
