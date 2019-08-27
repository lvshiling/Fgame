package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeBindGold, drop.DropResHandlerFunc(addBindGold))
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeGold, drop.DropResHandlerFunc(addGold))
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeSilver, drop.DropResHandlerFunc(addSilver))
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeExp, drop.DropResHandlerFunc(addExp))
}

func addBindGold(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.AddGold(resNum, true, reason, reasonText)

	return true
}

func addGold(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.AddGold(resNum, false, reason, reasonText)

	return true
}

func addSilver(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.AddSilver(resNum, reason, reasonText)

	return true
}

func addExp(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	if resNum <= 0 {
		return false
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.AddExp(resNum, reason, reasonText)

	return true
}
