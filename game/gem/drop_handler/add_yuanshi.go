package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	"fgame/fgame/game/gem/pbutil"
	playergem "fgame/fgame/game/gem/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeStorage, drop.DropResHandlerFunc(addYuanShi))
}

func addYuanShi(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	mine := manager.DropYuanShi(int32(resNum))
	scMsg := pbutil.BuildSCGemMineGet(mine)
	pl.SendMsg(scMsg)

	return true
}
