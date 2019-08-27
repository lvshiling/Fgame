package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	playerarena "fgame/fgame/game/arena/player"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeArenaPoint, drop.DropResHandlerFunc(addRes))
}

func addRes(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaManager.AddJiFen(int32(resNum))
	return true
}
