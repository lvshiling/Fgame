package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeYaoPai, drop.DropResHandlerFunc(addYaoPai))
}

func addYaoPai(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	manager.AddYaoPai(int32(resNum), reason, reasonText)

	return true
}
