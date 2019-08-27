package check

import (
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeField, wardrobe.CheckHandlerFunc(handleField))
}

func handleField(pl player.Player, seqId int32) (flag bool) {
	log.Debug("lingyu:处理判断是否幻化")
	manager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	return manager.IsUnrealed(int(seqId))
}
