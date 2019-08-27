package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/soulruins/pbutil"
	playersoulruins "fgame/fgame/game/soulruins/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSoulRuinsClear, command.CommandHandlerFunc(handleSoulRuinsClear))
}

//挑战过的关卡置空
func handleSoulRuinsClear(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空帝陵遗迹")

	err = clearSoulRuins(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空帝陵遗迹,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空帝陵遗迹完成")
	return
}

//帝陵遗迹置为0
func clearSoulRuins(p scene.Player) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	manager.GMClearSoulRuins()

	scSoulRuinsGet := pbutil.BuildSCSoulRuinsGet(pl)
	pl.SendMsg(scSoulRuinsGet)
	return
}
