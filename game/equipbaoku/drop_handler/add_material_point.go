package drop_handler

import (
	commonlogtypes "fgame/fgame/common/log/types"
	commontypes "fgame/fgame/game/common/types"
	"fgame/fgame/game/drop/drop"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equiptypes "fgame/fgame/game/equipbaoku/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	drop.RegistDropResHandler(itemtypes.ItemAutoUseResSubTypeMaterialBaoKuAttendPoints, drop.DropResHandlerFunc(addMaterialPoint))
}

func addMaterialPoint(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	manager := pl.GetPlayerDataManager(playertypes.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	return manager.AttendEquipBaoKu(0, int32(resNum), 0, commontypes.ChangeTypeItemGet, equiptypes.BaoKuTypeMaterials)
}
