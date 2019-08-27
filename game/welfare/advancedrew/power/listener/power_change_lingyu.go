package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	"fgame/fgame/game/player"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家领域战力变化
func playerLingyuPowerChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	power, ok := data.(int64)
	if !ok {
		return
	}

	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, power, welfaretypes.AdvancedTypeLingyu)
	return
}

func init() {
	gameevent.AddEventListener(lingyueventtypes.EventTypeLingyuPowerChanged, event.EventListenerFunc(playerLingyuPowerChange))
}
