package handler

import (
	"fgame/fgame/common/lang"
	feedbackfeelogic "fgame/fgame/game/feedbackfee/logic"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeExchangeMoney, command.CommandHandlerFunc(handleExchangeMoney))
}

func handleExchangeMoney(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	moneyStr := c.Args[0]
	money, err := strconv.ParseInt(moneyStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"moneyStr": moneyStr,
				"error":    err,
			}).Warn("gm:处理兑换,money不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	if money <= 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"money": money,
			}).Warn("gm:处理兑换,参数错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	feedbackfeelogic.HandlePlayerExchange(pl, int32(money))
	return
}
