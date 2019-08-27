package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playershenqi "fgame/fgame/game/shenqi/player"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeLingQi, drop.DropResHandlerFunc(addLingQi))
}

func addLingQi(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	return manager.AddLingQiNum(int64(resNum))
}
