package handler

import (
	"fgame/fgame/common/lang"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeArenapvpJiFen, command.CommandHandlerFunc(handleArenaJiFen))
}

func handleArenaJiFen(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	pointStr := c.Args[0]
	point, err := strconv.ParseInt(pointStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"point": point,
				"error": err,
			}).Warn("gm:处理比武大会积分,类型point不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	manager.GMSetJiFen(int32(point))
	return
}
