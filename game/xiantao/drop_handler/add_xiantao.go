package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerxiantao "fgame/fgame/game/xiantao/player"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeQianNianXianTao, drop.DropResHandlerFunc(addQianNianXianTao))
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeBaiNianXianTao, drop.DropResHandlerFunc(addBaiNianXianTao))
}

func addQianNianXianTao(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	manager.AddHighPeachCount(int32(resNum))
	return true
}

func addBaiNianXianTao(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	manager.AddJuniorPeachCount(int32(resNum))
	return true
}
