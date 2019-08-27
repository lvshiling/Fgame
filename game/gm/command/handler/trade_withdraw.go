package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"

	tradelogic "fgame/fgame/game/trade/logic"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTradeWithdraw, command.CommandHandlerFunc(handleTradeWithdraw))
}

func handleTradeWithdraw(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理交易下架")
	tpl, ok := pl.(player.Player)
	if !ok {
		log.Warn("gm:处理交易下架,不是玩家")
		return
	}

	if len(c.Args) < 1 {
		log.Warn("gm:处理交易下架,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	tradeIdStr := c.Args[0]

	tradeId, err := strconv.ParseInt(tradeIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"tradeId": tradeIdStr,
			}).Warn("gm:处理交易下架,tradeId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	err = tradeWithdraw(tpl, tradeId)

	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"tradeId": tradeId,
		}).Debug("gm:处理交易下架完成")
	return
}

func tradeWithdraw(p player.Player, tradeId int64) error {
	err := tradelogic.TradeWithdraw(p, tradeId)
	if err != nil {
		return err
	}
	return nil
}
