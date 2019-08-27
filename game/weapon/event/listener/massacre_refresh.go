package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	massacreeventtypes "fgame/fgame/game/massacre/event/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	weaponlogic "fgame/fgame/game/weapon/logic"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
)

//玩家戮仙刃变化影响兵魂
func playerMassacreWeapon(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	weaponEventData := data.(*massacreeventtypes.PlayerMassacreWeaponEventData)
	weaponId := weaponEventData.GetWeaponId()
	action := weaponEventData.GetAction()
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	if action {
		flag := weaponManager.WeaponActiveTemp(weaponId)
		if !flag {
			return
		}
		//同步属性
		weaponlogic.WeaponPropertyChanged(pl)
		scWeaponActive := pbutil.BuildSCWeaponActive(weaponId)
		pl.SendMsg(scWeaponActive)
	} else {
		flag := weaponManager.WeaponRemoveTemp(weaponId)
		if !flag {
			return
		}
		//同步属性
		weaponlogic.WeaponPropertyChanged(pl)
		weaponWear := weaponManager.GetWeaponWear()
		weaponMap := weaponManager.GetAllWeapon()
		scWeaponGet := pbutil.BuildSCWeaponGet(weaponWear, weaponMap)
		pl.SendMsg(scWeaponGet)
	}
	return
}

func init() {
	gameevent.AddEventListener(massacreeventtypes.EventTypeMassacreWeapon, event.EventListenerFunc(playerMassacreWeapon))
}
