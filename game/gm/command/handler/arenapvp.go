package handler

import (
	activitylogic "fgame/fgame/game/activity/logic"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

func init() {
	command.Register(gmcommandtypes.CommandTypePvp, command.CommandHandlerFunc(handlePvp))

}

func handlePvp(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	return activitylogic.HandleActiveAttend(pl, activitytypes.ActivityTypeArenapvp, "")

}
