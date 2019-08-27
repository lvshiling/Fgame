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
	command.Register(gmcommandtypes.CommandTypeGuanZhi, command.CommandHandlerFunc(handleGuanZhi))
}

func handleGuanZhi(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	guanZhiStr := c.Args[0]
	guanZhiInt, err := strconv.ParseInt(guanZhiStr, 10, 64)
	if err != nil {
		return
	}
	guanZhi := chuangshitypes.ChuangShiGuanZhi(guanZhiInt)
	if !guanZhi.Valid() {
		return
	}
	pl.SetGuanZhi(guanZhi)
	return
}
