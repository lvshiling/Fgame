package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playershenyu "fgame/fgame/game/shenyu/player"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeShenYuKey, drop.DropResHandlerFunc(addKey))
}

func addKey(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	shenYuManager := pl.GetPlayerDataManager(playertypes.PlayerShenYuDataManagerType).(*playershenyu.PlayerShenYuDataManager)
	shenYuManager.AddKeyNum(int32(resNum))

	return true
}
