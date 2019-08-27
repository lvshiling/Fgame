package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingyuventtypes "fgame/fgame/game/lingyu/event/types"
	"fgame/fgame/game/player"
	advancedrewrewlogic "fgame/fgame/game/welfare/advancedrew/rew/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家领域进阶
func playerLingyuAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeLingyu
	advancedrewrewlogic.UpdateAdvancedRewData(pl, advanceId, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(lingyuventtypes.EventTypeLingyuAdvanced, event.EventListenerFunc(playerLingyuAdavanced))
}
