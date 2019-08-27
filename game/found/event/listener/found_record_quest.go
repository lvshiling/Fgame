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
	questtypes "fgame/fgame/game/quest/types"
)

//记录屠魔接受事件、排除天机牌
func questAcceptRecord(target event.EventTarget, data event.EventData) (err error) {
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
	if questtypes.QuestTypeTianJiPai == questType {
		return
	}

	resType, flag := foundtypes.QuestTypeToFoundResType(questType)
	if !flag {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	manager.IncreFoundResJoinTimes(resType)
	return
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestAccept, event.EventListenerFunc(questAcceptRecord))

}
