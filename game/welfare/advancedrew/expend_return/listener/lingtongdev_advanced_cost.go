package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/player"
	advancedrewexpendreturnlogic "fgame/fgame/game/welfare/advancedrew/expend_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//灵童系统进阶消耗
func lingTongDevAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	devObj, ok := data.(*playerlingtongdev.PlayerLingTongDevObject)
	if !ok {
		return
	}
	nexTemp := mount.GetMountService().GetMountNumber(devObj.GetAdvancedId() + 1)
	itemNum := nexTemp.ItemCount
	advancedType, ok := welfaretypes.LingTongDevTypeToAdvancedType(devObj.GetClassType())
	if !ok {
		return
	}

	//消耗返还（新版）
	advancedrewexpendreturnlogic.AdvancedExpendReturn(pl, itemNum, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevAdvancedCost, event.EventListenerFunc(lingTongDevAdvancedCost))
}
