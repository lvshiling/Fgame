package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//任务过5点
func questDailyCrossFive(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	dailyTag := data.(questtypes.QuestDailyTag)
	if !ok {
		return
	}
	//日环任务
	questlogic.GetNextDailyQuest(pl, dailyTag)
	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestDailyCrossFive, event.EventListenerFunc(questDailyCrossFive))
}
