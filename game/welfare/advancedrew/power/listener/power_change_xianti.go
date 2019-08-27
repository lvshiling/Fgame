package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
)

//玩家仙体战力变化
func playerXianTiPowerChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	power, ok := data.(int64)
	if !ok {
		return
	}

	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, power, welfaretypes.AdvancedTypeXianTi)
	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiPowerChanged, event.EventListenerFunc(playerXianTiPowerChange))
}
