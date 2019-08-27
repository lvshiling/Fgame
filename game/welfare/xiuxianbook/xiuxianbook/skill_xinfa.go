package xiuxianbook

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerskill "fgame/fgame/game/skill/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	RegisterCountLevelHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillXinFa, CountLevelHandlerFunc(countSkillXinfa))
}

func countSkillXinfa(pl player.Player) (level int32, err error) {
	// manager := pl.GetPlayerDataManager(playertypes.PlayerXinFaDataManagerType).(*playerxinfa.PlayerXinFaDataManager)
	// level = manager.GetAllXinfaLevel()
	manager := pl.GetPlayerDataManager(playertypes.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	skillMap := manager.GetRoleSkillMap()
	level = int32(0)
	for _, obj := range skillMap {
		for _, lev := range obj.TianFuMap {
			level += lev
		}
	}
	return level, nil
}
