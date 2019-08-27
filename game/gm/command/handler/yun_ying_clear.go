package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	playerwelfare "fgame/fgame/game/welfare/player"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeYunYingClear, command.CommandHandlerFunc(handleYunYingClear))
}

func handleYunYingClear(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) != 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	groupStr := c.Args[0]
	groupId, err := strconv.ParseInt(groupStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理运营活动重置,groupId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	//修改等级
	welfareManager := pl.GetPlayerDataManager(types.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.GMResetActivity(int32(groupId))
	return
}
