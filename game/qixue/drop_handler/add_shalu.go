package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/qixue/pbutil"
	playerqixue "fgame/fgame/game/qixue/player"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeShaLuXin, drop.DropResHandlerFunc(addShaLu))
}

func addShaLu(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
	manager.AddShaLu(int32(resNum))
	scMsg := pbutil.BuildSCQiXueShaQiVary(resNum)
	pl.SendMsg(scMsg)

	return true
}
