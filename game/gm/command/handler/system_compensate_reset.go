package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSystemCompensateReset, command.CommandHandlerFunc(handleSystemCompensateReset))

}

func handleSystemCompensateReset(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	pl.GMSetSystemCompensate(false)
	return
}
