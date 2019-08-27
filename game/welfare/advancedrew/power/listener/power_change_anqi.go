package listener

import (
	"fgame/fgame/core/event"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家暗器战力变化
func playerAnQiPowerChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	power, ok := data.(int64)
	if !ok {
		return
	}

	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, power, welfaretypes.AdvancedTypeAnqi)
	return
}

func init() {
	gameevent.AddEventListener(anqieventtypes.EventTypeAnqiPowerChanged, event.EventListenerFunc(playerAnQiPowerChange))
}
