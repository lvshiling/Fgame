package xiuxianbook

import (
	playergoldequip "fgame/fgame/game/goldequip/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	RegisterCountLevelHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight, CountLevelHandlerFunc(countEquipOpenlight))
}

func countEquipOpenlight(pl player.Player) (level int32, err error) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()
	level = goldequipBag.GetAllOpenlightLevel()
	return level, nil
}
