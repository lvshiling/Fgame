package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cache/cache"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

// 玩家队伍变化
func playerTeamChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	cache.GetCacheService().UpdateCache(player.ConvertFromPlayer(pl))
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerTeamChanged, event.EventListenerFunc(playerTeamChanged))
}
