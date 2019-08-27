package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	playertrade "fgame/fgame/game/trade/player"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypePlayerTradeRecycle, command.CommandHandlerFunc(handlePlayerTradeRecycle))

}

func handlePlayerTradeRecycle(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	recycleStr := c.Args[0]
	recycle, err := strconv.ParseInt(recycleStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"recycle": recycleStr,
				"error":   err,
			}).Warn("gm:处理设置个人回收值,recycle不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if recycle < 0 {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"recycle": recycleStr,
				"error":   err,
			}).Warn("gm:处理设置个人回收值,recycle小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playertrade.PlayerTradeManager)
	manager.GMSetRecycle(recycle)
	return
}
