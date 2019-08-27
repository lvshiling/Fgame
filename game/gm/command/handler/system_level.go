package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeAllSystemLevel, command.CommandHandlerFunc(handleAllSystemLevel))

}

func handleAllSystemLevel(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理设置系统阶别,xianTiLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	handleMountLevel(p, c)
	handleWingLevel(p, c)
	handleAnqiLevel(p, c)
	handleFaBaoLevel(p, c)
	handleXianTiLevel(p, c)
	handleLingyuLevel(p, c)
	handleShenfaLevel(p, c)
	handleShiHunFanLevel(p, c)
	handleTianMoLevel(p, c)
	handleBodyShieldLevel(p, c)
	handleFeatherLevel(p, c)
	handleShieldLevel(p, c)

	for min := lingtongdevtypes.LingTongDevSysTypeMin; min <= lingtongdevtypes.LingTongDevSysTypeMax; min++ {
		lingTongDevAdvanced(pl, level, min)
	}
	return
}
