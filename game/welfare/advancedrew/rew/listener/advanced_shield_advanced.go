package listener

import (
	"fgame/fgame/core/event"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewrewlogic "fgame/fgame/game/welfare/advancedrew/rew/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家神盾尖刺进阶
func playerShieldAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeShield
	advancedrewrewlogic.UpdateAdvancedRewData(pl, advanceId, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeShieldAdvanced, event.EventListenerFunc(playerShieldAdavanced))
}
