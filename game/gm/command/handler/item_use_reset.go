package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeItemUseReset, command.CommandHandlerFunc(handleItemUseReset))
}

func handleItemUseReset(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理物品使用次数重置")

	err = itemUseReset(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理物品使用次数重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理物品使用次数重置完成")
	return
}

func itemUseReset(p scene.Player) (err error) {
	pl := p.(player.Player)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	inventoryManager.GMResetTimes()
	return
}
