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
	playerwing "fgame/fgame/game/wing/player"
)

//玩家属性变化
func playerPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	propertyEffectorType, ok := data.(playerpropertytypes.PropertyEffectorType)
	if !ok {
		return
	}
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeWing,
		playerpropertytypes.PlayerPropertyEffectorTypeFeather:
	default:
		return
	}
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeWing:
		{
			power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeWing)
			manager.WingPower(power)
			break
		}
	case playerpropertytypes.PlayerPropertyEffectorTypeFeather:
		{
			power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeFeather)
			manager.WingFeatherPower(power)
			break
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(playerPropertyChange))
}
