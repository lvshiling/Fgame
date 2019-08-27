package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfound "fgame/fgame/game/found/player"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
)

//日环任务事件
func questCommit(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	questId, ok := data.(int32)
	if !ok {
		return
	}

	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}

	questType := questTemplate.GetQuestType()
	resType, flag := foundtypes.QuestTypeToFoundResType(questType)
	if !flag {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	manager.IncreFoundResJoinTimes(resType)
	return
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestCommit, event.EventListenerFunc(questCommit))

}
