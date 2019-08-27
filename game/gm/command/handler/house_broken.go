package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerhouse "fgame/fgame/game/house/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeHouseBroken, command.CommandHandlerFunc(handleHouseBroken))

}

func handleHouseBroken(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	indexStr := c.Args[0]
	houseIndex, err := strconv.ParseInt(indexStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"houseIndex": houseIndex,
				"error":      err,
			}).Warn("gm:处理设置房子损坏,房子序号不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	houseManager := pl.GetPlayerDataManager(playertypes.PlayerHouseDataManagerType).(*playerhouse.PlayerHouseDataManager)
	houseManager.GmSetHouseBroken(int32(houseIndex))

	return
}
