package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeItemClear, command.CommandHandlerFunc(handleItemClear))
}

func handleItemClear(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空物品")

	err = clearAll(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空物品,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空物品完成")
	return
}

func clearAll(p scene.Player) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	manager.GMClearAll()
	logic.SnapInventoryChanged(pl)
	return
}
