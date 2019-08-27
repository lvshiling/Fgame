package xiuxianbook

import (
	playerdianxing "fgame/fgame/game/dianxing/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	RegisterCountLevelHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeDianXing, CountLevelHandlerFunc(countDianxing))
}

func countDianxing(pl player.Player) (level int32, err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	level = manager.GetDianXingObject().CurrLevel
	return level, nil
}
