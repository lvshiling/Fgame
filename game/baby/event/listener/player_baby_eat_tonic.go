package listener

import (
	"fgame/fgame/core/event"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	"fgame/fgame/game/cache/cache"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

// 玩家宝宝怀孕信息变化
func playerBabyPregnantChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	cache.GetCacheService().UpdateCache(player.ConvertFromPlayer(pl))
	return
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyPregnantChanged, event.EventListenerFunc(playerBabyPregnantChanged))
}
