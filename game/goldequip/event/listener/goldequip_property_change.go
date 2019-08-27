package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playergoldequip "fgame/fgame/game/goldequip/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(goldEquipPropertyChange))
}

func goldEquipPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	propertyEffectorType := data.(playerpropertytypes.PropertyEffectorType)
	if propertyEffectorType != playerpropertytypes.PlayerPropertyEffectorTypeGoldequip {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power := propertyManager.GetModuleForce(propertyEffectorType)

	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldequipManager.SetPower(power)

	return
}
