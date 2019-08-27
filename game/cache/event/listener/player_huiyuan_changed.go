package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/cache/cache"
	gameevent "fgame/fgame/game/event"
	huiyuaneventtypes "fgame/fgame/game/huiyuan/event/types"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
)

// 玩家会员变化
func playerHuiYuanChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	huiyuanType, ok := data.(huiyuantypes.HuiYuanType)
	if !ok {
		return
	}

	if huiyuanType != huiyuantypes.HuiYuanTypePlus {
		return
	}

	pl.SyncHuiYuan(true)
	cache.GetCacheService().UpdateCache(player.ConvertFromPlayer(pl))
	return
}

func init() {
	gameevent.AddEventListener(huiyuaneventtypes.EventTypeHuiYuanBuy, event.EventListenerFunc(playerHuiYuanChanged))
}
