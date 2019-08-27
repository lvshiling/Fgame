package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	playerchuangshi "fgame/fgame/game/chuangshi/player"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeChuangShiJifen, drop.DropResHandlerFunc(addAutoRes))
}

func addAutoRes(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	chuangshiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
	chuangshiManager.AddJiFen(resNum)

	return true
}
