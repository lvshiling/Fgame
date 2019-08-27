package xiuxianbook

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playersoul "fgame/fgame/game/soul/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	RegisterCountLevelHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillDiHun, CountLevelHandlerFunc(countSkillDihun))
}

func countSkillDihun(pl player.Player) (level int32, err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	level = manager.GetAllSoulSkillLevel()
	return level, nil
}
