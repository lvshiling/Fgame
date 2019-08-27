package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/scene/scene"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeJieYi, command.CommandHandlerFunc(handleJieYi))
}

func handleJieYi(p scene.Player, c *command.Command) (err error) {
	// pl := p.(player.Player)
	// if len(c.Args) < 1 {
	// 	pl.SetJieYiName("")
	// 	return
	// }

	// name := c.Args[0]
	// pl.SetJieYiName(name)
	// return
	return
}
