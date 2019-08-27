package listener

import (
	"fgame/fgame/core/event"
	anqiventtypes "fgame/fgame/game/anqi/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedblessfeedbacklogic "fgame/fgame/game/welfare/advanced/bless_feedback/logic"
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
	advancedType := welfaretypes.AdvancedTypeAnqi

	// 进阶奖励（每天类型变更，顺序随AdvancedType）
	advancedblessfeedbacklogic.UpdateAdvancedBlessActivityData(pl, advanceId, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(anqiventtypes.EventTypeAnqiAdvanced, event.EventListenerFunc(playerAnqiAdavanced))
}
