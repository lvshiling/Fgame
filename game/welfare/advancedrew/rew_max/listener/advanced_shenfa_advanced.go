package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	advancedrewrewlogic "fgame/fgame/game/welfare/advancedrew/rew_max/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家身法进阶
func playerShenfaAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}
	advancedType := welfaretypes.AdvancedTypeShenfa
	advancedrewrewlogic.UpdateAdvancedRewData(pl, advanceId, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaAdvanced, event.EventListenerFunc(playerShenfaAdavanced))
}
