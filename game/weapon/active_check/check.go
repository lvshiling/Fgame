package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"
	playerweapon "fgame/fgame/game/weapon/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeWeapon, wardrobe.CheckHandlerFunc(handleWeapon))
}

func handleWeapon(pl player.Player, seqId int32) (flag bool) {
	log.Debug("weapon:处理判断是否已激活")
	manager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	weaponInfo := manager.GetWeapon(seqId)
	if weaponInfo == nil {
		return
	}
	return weaponInfo.ActiveFlag == 1
}
