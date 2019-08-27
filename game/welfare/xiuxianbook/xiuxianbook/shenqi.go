package xiuxianbook

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playershenqi "fgame/fgame/game/shenqi/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	RegisterCountLevelHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeShenQi, CountLevelHandlerFunc(countShenqi))
}

func countShenqi(pl player.Player) (level int32, err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	level = manager.GetAllShenQiLevel()
	return level, nil
}
