package listener

import (
	"fgame/fgame/core/event"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(additionSysPropertyChange))
}

func additionSysPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	propertyEffectorType := data.(playerpropertytypes.PropertyEffectorType)
	if propertyEffectorType != playerpropertytypes.PlayerPropertyEffectorTypeShengHen {
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power := propertyManager.GetModuleForce(propertyEffectorType)

	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	additionsysManager.SetShengHenPower(power)

	return
}
