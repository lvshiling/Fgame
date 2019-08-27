package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	towereventtypes "fgame/fgame/game/tower/event/types"
)

//打宝消耗时间
func toverDaBaoNotice(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	daBaoStarTime, ok := data.(int64)
	if !ok {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	daBaoTime := int32(now - daBaoStarTime)

	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeDaBaoTimePast, 0, daBaoTime)
	return
}

func init() {
	gameevent.AddEventListener(towereventtypes.EventTypeTowerDaBaoNotice, event.EventListenerFunc(toverDaBaoNotice))
}
