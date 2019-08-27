package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家身法战力变化
func playerShenfaPowerChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	power, ok := data.(int64)
	if !ok {
		return
	}

	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, power, welfaretypes.AdvancedTypeShenfa)
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaPowerChanged, event.EventListenerFunc(playerShenfaPowerChange))
}
