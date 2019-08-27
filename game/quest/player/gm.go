package player

import (
	"fgame/fgame/game/global"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
)

//以下函数仅GM命令使用
func (pqdm *PlayerQuestDataManager) gmInitQuest(questId int32) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	qu := pqdm.GetQuestById(questId)
	if qu == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	oldState := qu.QuestState
	qu.QuestState = questtypes.QuestStateInit
	qu.CollectItemDataMap = make(map[int32]int32)
	qu.QuestDataMap = make(map[int32]int32)
	qu.UpdateTime = now
	qu.SetModified()
	switch oldState {
	case questtypes.QuestStateInit:
		break
	default:
		pqdm.removeQuestByIdAndState(oldState, questId)
		pqdm.addQuest(qu)
		break
	}
	for _, nextQuestId := range questTemplate.GetNextQuestIds() {
		pqdm.gmInitQuest(nextQuestId)
	}
	return
}

func (pqdm *PlayerQuestDataManager) gmCommitQuest(questId int32) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	qu := pqdm.GetQuestById(questId)
	if qu == nil {
		pqdm.AddQuest(questId)
	}
	qu = pqdm.GetQuestById(questId)
	if qu.QuestState == questtypes.QuestStateCommit {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	oldState := qu.QuestState
	qu.QuestState = questtypes.QuestStateCommit
	qu.UpdateTime = now
	qu.SetModified()
	pqdm.removeQuestByIdAndState(oldState, questId)
	pqdm.addQuest(qu)
	for _, prevQuestId := range questTemplate.GetPrevQuestIds() {
		pqdm.gmCommitQuest(prevQuestId)
	}
	return
}

//gm修改任务
func (pqdm *PlayerQuestDataManager) GMModifyQuestId(questId int32) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	for _, nextQuestId := range questTemplate.GetNextQuestIds() {
		pqdm.gmInitQuest(nextQuestId)
	}
	for _, prevQuestId := range questTemplate.GetPrevQuestIds() {
		pqdm.gmCommitQuest(prevQuestId)
	}
	qu := pqdm.GetQuestById(questId)
	if qu == nil {
		pqdm.AddQuest(questId)

		pqdm.ActiveQuest(questId)

		pqdm.AcceptQuest(questId)

	} else {
		now := global.GetGame().GetTimeService().Now()
		oldState := qu.QuestState
		qu.QuestState = questtypes.QuestStateAccept
		qu.UpdateTime = now
		qu.SetModified()
		pqdm.removeQuestByIdAndState(oldState, questId)
		pqdm.addQuest(qu)
	}
	return
}

//完成屠魔任务列表
func (pqdm *PlayerQuestDataManager) GMFinishTuMoTar() {
	now := global.GetGame().GetTimeService().Now()
	for _, levelQuestMap := range pqdm.tuMoQuestMap {
		for questId, obj := range levelQuestMap {
			if obj.QuestState != questtypes.QuestStateAccept {
				continue
			}
			q := pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
			//移除
			pqdm.removeQuestByIdAndState(questtypes.QuestStateAccept, questId)
			q.QuestState = questtypes.QuestStateFinish
			q.UpdateTime = now
			q.SetModified()

			pqdm.addQuest(q)
		}
	}
	return
}

//清空屠魔次数
func (pqdm *PlayerQuestDataManager) GMClearTuMoNum() {
	now := global.GetGame().GetTimeService().Now()
	pqdm.playerTuMoObject.Num = 0
	pqdm.playerTuMoObject.LastTime = 0
	pqdm.playerTuMoObject.UpdateTime = now
	pqdm.playerTuMoObject.SetModified()
}

//日环任务设置次数
func (pqdm *PlayerQuestDataManager) GMSetDailyQuestTimes(times questtypes.QuestDailyType) (questId int32, quest *PlayerQuestObject) {

	return
}

//设置指定 日环任务
func (pqdm *PlayerQuestDataManager) GMSetDailyQuest(dailyTag questtypes.QuestDailyTag, seqId int32) (updateQuestList []*PlayerQuestObject) {
	dailyTempalte := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, seqId)
	if dailyTempalte == nil {
		return
	}

	var questList []*PlayerQuestObject
	personDailyList, allianceDailyList := pqdm.getQuestTag()
	for questDailyTag, _ := range pqdm.playerDailyObjectMap {
		if questDailyTag != dailyTag {
			continue
		}

		switch questDailyTag {
		case questtypes.QuestDailyTagPerson:
			questList = personDailyList
		case questtypes.QuestDailyTagAlliance:
			questList = allianceDailyList
		}
		updateQuestList = pqdm.resetDaily(questDailyTag, questList)
	}

	quest := pqdm.AddDailyQuest(dailyTag, dailyTempalte)
	updateQuestList = append(updateQuestList, quest)
	return
}
