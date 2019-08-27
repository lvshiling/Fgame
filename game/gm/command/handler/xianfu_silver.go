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
	command.Register(gmcommandtypes.CommandTypeXianfuSilver, command.CommandHandlerFunc(handleXianfuSilver))
}

func handleXianfuSilver(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理进入秘境仙府银两副本")
	pl := p.(player.Player)
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
			}).Warn("gm:处理跳转秘境仙府银两副本,秘境仙府副本id不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	err = enterXianfuSilver(pl, int32(level))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": level,
				"error": err,
			}).Warn("gm:处理跳转秘境仙府银两副本,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":    pl.GetId(),
			"level": level,
		}).Debug("gm:处理跳转秘境仙府银两副本,完成")
	return
}

func enterXianfuSilver(pl player.Player, level int32) (err error) {
	xianfulogic.PlayerEnterXianFu(pl, xianfutypes.XianfuTypeSilver, level)
	return
}
