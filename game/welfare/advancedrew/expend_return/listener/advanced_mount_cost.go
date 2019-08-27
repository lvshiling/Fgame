package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/player"
	advancedrewexpendreturnlogic "fgame/fgame/game/welfare/advancedrew/expend_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//坐骑进阶消耗
func playerMountAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}
	nexTemp := mount.GetMountService().GetMountNumber(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeMount

	//消耗返还（新版）
	advancedrewexpendreturnlogic.AdvancedExpendReturn(pl, itemNum, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountAdvancedCost, event.EventListenerFunc(playerMountAdvancedCost))
}
