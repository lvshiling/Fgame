package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/drop/drop"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/massacre/pbutil"
	playermassacre "fgame/fgame/game/massacre/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeShaQi, drop.DropResHandlerFunc(addShaQi))
}

func addShaQi(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	_ = manager.DropShaQi(int32(resNum))
	scMsg := pbutil.BuildSCMassacreShaQiVary(resNum)
	pl.SendMsg(scMsg)

	return true
}
