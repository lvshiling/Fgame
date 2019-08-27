package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	luckyeventtypes "fgame/fgame/game/lucky/event/types"
	playerlucky "fgame/fgame/game/lucky/player"
	"fgame/fgame/game/player"
)

// 添加幸运类型
func playerLuckyAdd(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	luckyInfo, ok := data.(*playerlucky.PlayerLuckyObject)
	if !ok {
		return
	}

	// 同步幸运信息
	itemId := luckyInfo.GetItemId()
	itemTemp := item.GetItemService().GetItem(int(itemId))
	typ := itemTemp.GetItemType()
	subType := itemTemp.GetItemSubType()
	rate := itemTemp.TypeFlag1
	expireTime := luckyInfo.GetExpireTime()
	pl.AddLucky(typ, subType, rate, expireTime)
	return
}

func init() {
	gameevent.AddEventListener(luckyeventtypes.EventTypeLuckyAdd, event.EventListenerFunc(playerLuckyAdd))
}
