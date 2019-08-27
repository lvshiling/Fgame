package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
)

//接取天机牌任务
func (pqdm *PlayerQuestDataManager) AcceptSecretCardQuest(questId int32) (questObj *PlayerQuestObject, flag bool) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	typ := questTemplate.GetQuestType()
	if typ != questtypes.QuestTypeTianJiPai {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	questObj = pqdm.GetQuestById(questId)
	if questObj == nil {
		questObj = createQuest(pqdm.p, questId)
	} else {
		if questObj.QuestState != questtypes.QuestStateCommit &&
			questObj.QuestState != questtypes.QuestStateDiscard {
			return
		}
		questObj.QuestDataMap = make(map[int32]int32)
		questObj.CollectItemDataMap = make(map[int32]int32)
		//移除
		pqdm.removeQuestByIdAndState(questObj.QuestState, questId)
	}
	questObj.UpdateTime = now
	questObj.QuestState = questtypes.QuestStateAccept
	questObj.SetModified()

	pqdm.addQuest(questObj)
	gameevent.Emit(questeventtypes.EventTypeQuestAccept, pqdm.p, questId)
	flag = true
	return
}
