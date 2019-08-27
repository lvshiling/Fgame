package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playershenmo "fgame/fgame/game/shenmo/player"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeGongXun, drop.DropResHandlerFunc(addGongXun))
}

func addGongXun(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerShenMoWarDataManagerType).(*playershenmo.PlayerShenMoDataManager)
	return manager.AddGongXun(int32(resNum))
}
