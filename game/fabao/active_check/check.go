package check

import (
	playerfabao "fgame/fgame/game/fabao/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeFaBao, wardrobe.CheckHandlerFunc(handleFaBao))
}

func handleFaBao(pl player.Player, seqId int32) (flag bool) {
	log.Debug("fabao:处理判断是否幻化")
	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	return manager.IsUnrealed(int(seqId))
}
