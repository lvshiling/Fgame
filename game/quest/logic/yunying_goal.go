package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
)

// 是否可以激活运营活动目标任务
func InitYunYingGoalQuest(pl player.Player, goalId int32) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)

	// 子任务
	var questList []*playerquest.PlayerQuestObject
	goalTemp := questtemplate.GetQuestTemplateService().GetYunYingGoalTemplate(goalId)
	for questId, _ := range goalTemp.GetQuestMap() {
		qu := manager.GetQuestById(questId)
		if qu != nil {
			continue
		}

		manager.AddQuest(questId)
		questList = append(questList, CheckInitQuestList(pl)...)
	}

	if len(questList) > 0 {
		scMsgUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scMsgUpdate)
	}
}
