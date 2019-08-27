package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"
	playerwing "fgame/fgame/game/wing/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeWing, wardrobe.CheckHandlerFunc(handleWing))
}

func handleWing(pl player.Player, seqId int32) (flag bool) {
	log.Debug("wing:处理判断是否已幻化")
	manager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	return manager.IsUnrealed(int(seqId))
}
