package listener

import (
	"fgame/fgame/core/event"
	anqiventtypes "fgame/fgame/game/anqi/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewrewextendedlogic "fgame/fgame/game/welfare/advancedrew/rew_extended/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家暗器进阶
func playerAnqiAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	// 进阶奖励(随功能开启)
	advancedType := welfaretypes.AdvancedTypeAnqi
	advancedrewrewextendedlogic.UpdateAdvancedRewExtendedData(pl, advanceId, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(anqiventtypes.EventTypeAnqiAdvanced, event.EventListenerFunc(playerAnqiAdavanced))
}
