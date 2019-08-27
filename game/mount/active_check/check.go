package check

import (
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeMount, wardrobe.CheckHandlerFunc(handleMount))
}

func handleMount(pl player.Player, seqId int32) (flag bool) {
	log.Debug("mount:处理判断是否激活")
	manager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	return manager.IsUnrealed(int(seqId))
}
