package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tianmotieventtypes "fgame/fgame/game/tianmo/event/types"
	advancedrewpowerlogic "fgame/fgame/game/welfare/advancedrew/power/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家天魔体战力变化
func playerTianMoTiFanPowerChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	power, ok := data.(int64)
	if !ok {
		return
	}

	advancedrewpowerlogic.UpdateAdvancedPowerData(pl, power, welfaretypes.AdvancedTypeTianMoTi)
	return
}

func init() {
	gameevent.AddEventListener(tianmotieventtypes.EventTypeTianMoPowerChanged, event.EventListenerFunc(playerTianMoTiFanPowerChange))
}
