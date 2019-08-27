package handler

import (
	"fgame/fgame/common/lang"
	playerbagua "fgame/fgame/game/bagua/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeBaGuaBuChang, command.CommandHandlerFunc(handleBaGuaBuChang))

}

func handleBaGuaBuChang(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm: 处理八卦补偿置零命令")
	pl := p.(player.Player)
	if len(c.Args) > 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	manager.GmSetBuChang()

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		},
	).Debug("gm: 处理八卦补偿置零命令成功")

	return
}
