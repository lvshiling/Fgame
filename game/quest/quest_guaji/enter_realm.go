package quest_guaji

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	realmlogic "fgame/fgame/game/realm/logic"
	gametemplate "fgame/fgame/game/template"
)

//进入天劫塔X次
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeEnterRealm, guaji.QuestGuaJiFunc(enterRealm))

}

//进入天劫塔X次
func enterRealm(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	num := q.QuestDataMap[0]
	needNum := demandMap[0]
	if num >= needNum {
		return true
	}
	realmlogic.HandleTianJieTa(p)
	return true
}
