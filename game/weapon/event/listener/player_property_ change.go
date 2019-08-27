package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playerweapon "fgame/fgame/game/weapon/player"
)

//玩家属性变化
func playerPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	propertyEffectorType := data.(playerpropertytypes.PropertyEffectorType)
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeWeapon:
	default:
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	basePower := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeWeapon)
	power := basePower
	manager.WeaponPower(power)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(playerPropertyChange))
}
