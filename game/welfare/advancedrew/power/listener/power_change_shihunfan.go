package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家噬魂幡战力变化
func playerShiHunFanPowerChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	power, ok := data.(int64)
	if !ok {
		return
	}

	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, power, welfaretypes.AdvancedTypeShiHunFan)
	return
}

func init() {
	gameevent.AddEventListener(shihunfaneventtypes.EventTypeShiHunFanPowerChanged, event.EventListenerFunc(playerShiHunFanPowerChange))
}
