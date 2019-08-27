package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	playerfourgod "fgame/fgame/game/fourgod/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeKey, drop.DropResHandlerFunc(addKey))
}

func addKey(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	manager.AddKeyNum(int32(resNum))

	return true
}
