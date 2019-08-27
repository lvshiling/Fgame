package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/onearena/onearena"
	"fgame/fgame/game/scene/scene"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeResetKunTime, command.CommandHandlerFunc(handleResetKunTime))
}

func handleResetKunTime(p scene.Player, c *command.Command) (err error) {
	onearena.GetOneArenaService().GmResetKunTime()
	return
}
