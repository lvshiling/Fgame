package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/welfare"
)

//添加疯狂宝箱日志
func addCrazyBoxLog(target event.EventTarget, data event.EventData) (err error) {
	groupId := target.(int32)
	eventData, ok := data.(*welfareeventtypes.CrazyBoxAddLogEventData)
	if !ok {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		return
	}

	itemId := eventData.GetItemId()
	plName := eventData.GetPlayerName()
	itemNum := eventData.GetItemNum()
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}
	quality := itemTemplate.GetQualityType()
	if quality < itemtypes.ItemQualityTypeOrange {
		return
	}
	// 添加日志
	welfare.GetWelfareService().AddCrazyBoxLog(groupId, plName, itemId, itemNum)

	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeCrazyBoxAddLog, event.EventListenerFunc(addCrazyBoxLog))
}
