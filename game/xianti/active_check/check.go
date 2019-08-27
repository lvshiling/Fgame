package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"
	playerxianti "fgame/fgame/game/xianti/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeXianTi, wardrobe.CheckHandlerFunc(handleXianTi))
}

func handleXianTi(pl player.Player, seqId int32) (flag bool) {
	log.Debug("xianti:处理判断是否幻化")
	manager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	return manager.IsUnrealed(int(seqId))
}
