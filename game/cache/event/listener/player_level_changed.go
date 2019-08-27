package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/cache/cache"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
)

// 玩家等级变化
func playerLevelChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	cache.GetCacheService().UpdateCache(player.ConvertFromPlayer(pl))
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerLevelChanged, event.EventListenerFunc(playerLevelChanged))
}
