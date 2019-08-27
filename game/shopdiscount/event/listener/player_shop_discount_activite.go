package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/shopdiscount/pbutil"
	playershopdiscount "fgame/fgame/game/shopdiscount/player"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
)

//玩家激活商城促销特权
func playerShopDiscountActivite(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventDate, ok := data.(*welfareeventtypes.PlayerShopDiscountEventData)
	if !ok {
		return
	}

	typ := eventDate.GetShopDiscountType()
	startTime := eventDate.GetStartTime()
	endTime := eventDate.GetEndTime()
	manager := pl.GetPlayerDataManager(types.PlayerShopDiscountDataManagerType).(*playershopdiscount.PlayerShopDiscountDataManager)
	obj := manager.GetShopDiscountObj()
	if obj.EndTime >= endTime {
		return
	}

	manager.SetCurShopDiscountType(typ, startTime, endTime)
	scMsg := pbutil.BuildSCShopDiscountNotice(obj)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeShopDiscountActivite, event.EventListenerFunc(playerShopDiscountActivite))
}
