package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

//玩家属性变化
func playerPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	propertyEffectorType := data.(playerpropertytypes.PropertyEffectorType)
	if propertyEffectorType != playerpropertytypes.PlayerPropertyEffectorTypeMount {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeMount)
	manager.MountPower(power)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(playerPropertyChange))
}
