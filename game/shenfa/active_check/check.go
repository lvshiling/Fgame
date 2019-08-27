package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playershenfa "fgame/fgame/game/shenfa/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeShenFa, wardrobe.CheckHandlerFunc(handleShenFa))
}

func handleShenFa(pl player.Player, seqId int32) (flag bool) {
	log.Debug("shenfa:处理判断是否幻化")
	manager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	return manager.IsUnrealed(int(seqId))
}
