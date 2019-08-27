package quest_guaji

import (
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
)

//灵童激活
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeLingTongActivateNum, guaji.QuestGuaJiFunc(lingTongActive))
}

//激活灵童
func lingTongActive(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	needNum := demandMap[0]
	num := q.QuestDataMap[0]
	if num >= needNum {
		return true
	}
	for _, lingTongTemplate := range lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplateList() {
		if lingtonglogic.CheckIfLingTongActivate(p, int32(lingTongTemplate.TemplateId())) {
			lingtonglogic.HandleLingTongActivate(p, int32(lingTongTemplate.TemplateId()))
			return true
		}
	}

	return false
}
