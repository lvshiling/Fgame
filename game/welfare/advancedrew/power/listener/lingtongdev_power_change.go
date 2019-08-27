package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/player"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//灵童系统战力变化
func lingTongDevPowerChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	devObj, ok := data.(*playerlingtongdev.PlayerLingTongPowerObject)
	if !ok {
		return
	}

	advancedType, ok := welfaretypes.LingTongDevTypeToAdvancedType(devObj.GetClassType())
	if !ok {
		return
	}
	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, devObj.GetPower(), advancedType)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevPowerChanged, event.EventListenerFunc(lingTongDevPowerChanged))
}
