package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	playershenyu "fgame/fgame/game/shenyu/player"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeShenYuKey, command.CommandHandlerFunc(handleShenYuKey))
}

func handleShenYuKey(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	keyStr := c.Args[0]
	shenYuKey, err := strconv.ParseInt(keyStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"shenYuKey": shenYuKey,
				"error":     err,
			}).Warn("gm:处理设置神域钥匙,shenYuKey不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerShenYuDataManagerType).(*playershenyu.PlayerShenYuDataManager)
	manager.GMSetKeyNum(int32(shenYuKey))
	return
}
