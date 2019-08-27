package listener

import (
	"fgame/fgame/core/event"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
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
	if propertyEffectorType != playerpropertytypes.PlayerPropertyEffectorTypeBodyShield &&
		propertyEffectorType != playerpropertytypes.PlayerPropertyEffectorTypeShield {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeBodyShield:
		{
			power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeBodyShield)
			manager.BodyShieldPower(power)
			break
		}
	case playerpropertytypes.PlayerPropertyEffectorTypeShield:
		{
			power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeShield)
			manager.ShieldPower(power)
			break
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(playerPropertyChange))
}
