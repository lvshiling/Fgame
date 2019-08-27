package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/dianxing/pbutil"
	playerdianxing "fgame/fgame/game/dianxing/player"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeXingChen, drop.DropResHandlerFunc(addXingChen))
}

func addXingChen(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string)bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	_ = manager.DropXingChen(int32(resNum))
	scMsg := pbutil.BuildSCDianxingXingchenVary(resNum)
	pl.SendMsg(scMsg)

	return true
}
