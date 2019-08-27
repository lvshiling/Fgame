package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeJieYiSetShengWeiZhi, command.CommandHandlerFunc(handleJieYiSetShengWeiZhi))
}

func handleJieYiSetShengWeiZhi(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置结义声威值")

	pl := p.(player.Player)

	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"numStr": numStr,
				"error":  err,
			}).Warn("gm:处理设置声威值,shengweizhi不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	jieYiManager.GmSetShengWeiZhi(int32(num))

	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理设置声威值,完成")
	return
}
