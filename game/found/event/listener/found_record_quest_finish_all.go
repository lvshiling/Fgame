package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfound "fgame/fgame/game/found/player"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questeventtypes "fgame/fgame/game/quest/event/types"
)

//记录天机牌、屠魔接受完成所有事件
func questFinishAll(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*questeventtypes.QuestFinishAllEventData)
	if !ok {
		return
	}

	questType := eventData.GetQuestType()
	num := eventData.GetNum()
	if num <= 0 {
		return
	}

	resType, flag := foundtypes.QuestTypeToFoundResType(questType)
	if !flag {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	manager.IncreFoundResJoinTimesBatch(resType, num)
	return
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestFinishAll, event.EventListenerFunc(questFinishAll))
}
