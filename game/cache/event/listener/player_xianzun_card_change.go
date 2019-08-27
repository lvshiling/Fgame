package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/cache/cache"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	xianzuncardeventtypes "fgame/fgame/game/xianzuncard/event/types"
)

// 玩家仙尊特权卡信息变化
func playerXianZunCardChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	cache.GetCacheService().UpdateCache(player.ConvertFromPlayer(pl))
	return
}

func init() {
	gameevent.AddEventListener(xianzuncardeventtypes.EventTypeXianZunCardDataChange, event.EventListenerFunc(playerXianZunCardChanged))
}
