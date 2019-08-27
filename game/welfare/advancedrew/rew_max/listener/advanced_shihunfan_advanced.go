package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shihunfanventtypes "fgame/fgame/game/shihunfan/event/types"
	advancedrewrewlogic "fgame/fgame/game/welfare/advancedrew/rew_max/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家噬魂幡进阶
func playerShiHunFanAdavancedRew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeShiHunFan
	advancedrewrewlogic.UpdateAdvancedRewData(pl, int32(advanceId), advancedType)
	return
}

func init() {
	gameevent.AddEventListener(shihunfanventtypes.EventTypeShiHunFanAdvanced, event.EventListenerFunc(playerShiHunFanAdavancedRew))
}