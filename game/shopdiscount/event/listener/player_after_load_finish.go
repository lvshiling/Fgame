package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/shopdiscount/pbutil"
	playershopdiscount "fgame/fgame/game/shopdiscount/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerShopDiscountDataManagerType).(*playershopdiscount.PlayerShopDiscountDataManager)
	obj := manager.GetShopDiscountObj()
	scShopDiscountGet := pbutil.BuildSCShopDiscountGet(obj)
	pl.SendMsg(scShopDiscountGet)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
