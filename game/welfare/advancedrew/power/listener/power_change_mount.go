package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	"fgame/fgame/game/player"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家坐骑战力变化
func playerMountPowerChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	power, ok := data.(int64)
	if !ok {
		return
	}

	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, power, welfaretypes.AdvancedTypeMount)
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountPowerChanged, event.EventListenerFunc(playerMountPowerChange))
}
