package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
)

//玩家仙羽战力变化
func playerFeatherPowerChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	power, ok := data.(int64)
	if !ok {
		return
	}

	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, power, welfaretypes.AdvancedTypeFeather)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeFeatherPowerChanged, event.EventListenerFunc(playerFeatherPowerChange))
}
