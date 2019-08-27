package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingyuventtypes "fgame/fgame/game/lingyu/event/types"
	"fgame/fgame/game/player"
	advancedrewrewextendedlogic "fgame/fgame/game/welfare/advancedrew/rew_extended/logic"
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
	advancedrewrewextendedlogic.UpdateAdvancedRewExtendedData(pl, advanceId, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(lingyuventtypes.EventTypeLingyuAdvanced, event.EventListenerFunc(playerLingyuAdavanced))
}
