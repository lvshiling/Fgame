package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	vipeventtypes "fgame/fgame/game/vip/event/types"
)

// vip升级或升星
func vipUplevel(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	// 更新vip属性
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeVipLevel.Mask())

	return
}

func init() {
	gameevent.AddEventListener(vipeventtypes.EventTypeVipLevelChanged, event.EventListenerFunc(vipUplevel))
}
