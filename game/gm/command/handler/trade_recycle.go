package handler

import (
	"fgame/fgame/common/lang"

	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"

	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/trade/trade"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTradeRecycle, command.CommandHandlerFunc(handleTradeRecycle))

}

func handleTradeRecycle(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理设置回收,recycle不是数字")
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
			}).Warn("gm:处理设置转生,recycle小于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	trade.GetTradeService().GMSetRecycle(recycle)
	return
}
