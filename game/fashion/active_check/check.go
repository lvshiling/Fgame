package check

import (
	playerfashion "fgame/fgame/game/fashion/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeFashion, wardrobe.CheckHandlerFunc(handleFashion))
}

func handleFashion(pl player.Player, seqId int32) (flag bool) {
	log.Debug("fashion:处理判断是否激活")
	manager := pl.GetPlayerDataManager(types.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	fashionInfo := manager.GetFashion(seqId)
	if fashionInfo == nil {
		return
	}
	flag = true
	return
}
