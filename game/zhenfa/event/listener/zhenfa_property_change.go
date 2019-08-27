package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playerzhenfa "fgame/fgame/game/zhenfa/player"
)

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(zhenFaPropertyChange))
}

func zhenFaPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	propertyEffectorType := data.(playerpropertytypes.PropertyEffectorType)
	if propertyEffectorType != playerpropertytypes.PlayerPropertyEffectorTypeZhenFa {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power := propertyManager.GetModuleForce(propertyEffectorType)

	zhenFaManager := pl.GetPlayerDataManager(playertypes.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	zhenFaManager.SetPower(power)

	return
}
