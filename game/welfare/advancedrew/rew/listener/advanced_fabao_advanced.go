package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fabaoventtypes "fgame/fgame/game/fabao/event/types"
	"fgame/fgame/game/player"
	advancedrewrewlogic "fgame/fgame/game/welfare/advancedrew/rew/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家法宝进阶
func playerFaBaoAdavancedRew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeFaBao
	advancedrewrewlogic.UpdateAdvancedRewData(pl, int32(advanceId), advancedType)

	return
}

func init() {
	gameevent.AddEventListener(fabaoventtypes.EventTypeFaBaoAdvanced, event.EventListenerFunc(playerFaBaoAdavancedRew))
}
