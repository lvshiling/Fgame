package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"

	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/trade/trade"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTradeRecycleTrigger, command.CommandHandlerFunc(handleTradeRecycleTrigger))

}

func handleTradeRecycleTrigger(p scene.Player, c *command.Command) (err error) {

	trade.GetTradeService().SystemRecycle()
	return
}
