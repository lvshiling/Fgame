package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeArenapvpJiFen, drop.DropResHandlerFunc(addRes))
}

func addRes(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpManager.AddJiFen(int32(resNum))
	return true
}
