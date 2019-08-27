package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

//玩家属性变化
func playerPropertyChange(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	propertyEffectorType := data.(playerpropertytypes.PropertyEffectorType)
	switch propertyEffectorType {
	case playerpropertytypes.PlayerPropertyEffectorTypeLingyu:
	default:
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingyu)
	manager.LingyuPower(power)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(playerPropertyChange))
}
