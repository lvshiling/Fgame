package xiuxianbook

import (
	playerlingtong "fgame/fgame/game/lingtong/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	RegisterCountLevelHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeLingTong, CountLevelHandlerFunc(countLingtong))
}

func countLingtong(pl player.Player) (level int32, err error) {
	lingtongManager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	level = lingtongManager.GetAllLingTongLevel()
	return level, nil
}
