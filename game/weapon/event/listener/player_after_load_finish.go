package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	weaponManager := pl.GetPlayerDataManager(playertypes.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	weaponWear := weaponManager.GetWeaponWear()
	weaponMap := weaponManager.GetAllWeapon()
	scWeaponGet := pbutil.BuildSCWeaponGet(weaponWear, weaponMap)
	pl.SendMsg(scWeaponGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
