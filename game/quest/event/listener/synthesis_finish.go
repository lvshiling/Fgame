package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	synthesiseventtypes "fgame/fgame/game/synthesis/event/types"
	synthesistypes "fgame/fgame/game/synthesis/types"
)

//合成完成
func synthesisFinsh(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*synthesiseventtypes.SynthesisFinishEventData)
	if !ok {
		return
	}
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ := eventData.GetType()
	num := eventData.GetNum()
	if num <= 0 {
		return
	}
	if typ != synthesistypes.TuMoLing {
		return
	}
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeSyntheticTuMoToken, 0, num)
	return
}

func init() {
	gameevent.AddEventListener(synthesiseventtypes.EventTypeSynthesisFinish, event.EventListenerFunc(synthesisFinsh))
}
