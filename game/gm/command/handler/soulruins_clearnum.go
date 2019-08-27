package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	playersoulruins "fgame/fgame/game/soulruins/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSoulRuinsClearNum, command.CommandHandlerFunc(handleSoulRuinsClearNum))
}

//挑战次数置为0
func handleSoulRuinsClearNum(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空帝陵遗迹挑战次数")

	err = clearSoulRuinsNum(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空帝陵遗迹挑战次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空帝陵遗迹挑战次数完成")
	return
}

//帝陵遗迹挑战次数置为0
func clearSoulRuinsNum(p scene.Player) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	manager.GMClearSoulRuinsNum()

	return nil
}
