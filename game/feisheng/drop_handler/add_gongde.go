package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeGongDe, drop.DropResHandlerFunc(addGongDe))
}

func addGongDe(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	manager.AddGongDe(resNum)

	return true
}
