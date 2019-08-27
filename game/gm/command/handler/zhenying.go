package handler

import (
	"fgame/fgame/common/lang"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeZhenYing, command.CommandHandlerFunc(handleZhenYing))
}

func handleZhenYing(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	campStr := c.Args[0]
	campInt, err := strconv.ParseInt(campStr, 10, 64)
	if err != nil {
		return
	}
	camp := chuangshitypes.ChuangShiCampType(campInt)
	if !camp.Valid() {
		return
	}
	pl.SetCamp(camp)
	return
}
