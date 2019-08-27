package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/feedbackfee/feedbackfee"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeExchangeFinish, command.CommandHandlerFunc(handleExchangeFinish))
}

func handleExchangeFinish(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	code := c.Args[0]
	if len(code) <= 0 {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:处理活动,code是空")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	feedbackfee.GetFeedbackFeeService().CodeExchangeByCode(code)
	return
}
