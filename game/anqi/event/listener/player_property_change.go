package listener

import (
	"fgame/fgame/core/event"
	playeranqi "fgame/fgame/game/anqi/player"
	gameevent "fgame/fgame/game/event"
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
	if propertyEffectorType != playerpropertytypes.PlayerPropertyEffectorTypeAnqi {
		return
	}

	anqiManager := pl.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeAnqi)
	anqiManager.AnqiPower(power)

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(playerPropertyChange))
}
