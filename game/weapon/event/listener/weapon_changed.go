package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	weaponeventtypes "fgame/fgame/game/weapon/event/types"
	playerweapon "fgame/fgame/game/weapon/player"
)

//兵魂变化
func weaponChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	weaponManager := pl.GetPlayerDataManager(playertypes.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	weaponId := weaponManager.GetWeaponWear()
	weaponState := weaponManager.GetWeaponState(weaponId)
	pl.SetWeapon(weaponId, int32(weaponState))
	return
}

func init() {
	gameevent.AddEventListener(weaponeventtypes.EventTypeWeaponChanged, event.EventListenerFunc(weaponChanged))
}
