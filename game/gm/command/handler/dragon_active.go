package handler

import (
	"fgame/fgame/common/lang"
	dragonlogic "fgame/fgame/game/dragon/logic"
	"fgame/fgame/game/dragon/pbutil"
	playerdragon "fgame/fgame/game/dragon/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeDragonActive, command.CommandHandlerFunc(handleDragonStage))

}

func handleDragonActive(p scene.Player, c *command.Command) (err error) {

	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerDragonDataManagerType).(*playerdragon.PlayerDragonDataManager)
	manager.GmDragonActive()

	dragonlogic.DragonPropertyChanged(pl)

	dragonInfo := manager.GetDragon()
	scDragonGet := pbuitl.BuildSCDragonGet(dragonInfo)
	pl.SendMsg(scDragonGet)
	return
}
