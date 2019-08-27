package commitflow

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCommitFlow(questtypes.QuestTypeQiYu, quest.CommitFlowHandlerFunc(handleCommitFlowQiYu))
}

//处理奇遇
func handleCommitFlowQiYu(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理奇遇")
	if questTemplate.GetQuestType() != questtypes.QuestTypeQiYu {
		return
	}
	questId := int32(questTemplate.TemplateId())
	qiyuId := questtemplate.GetQuestTemplateService().GetQiYuIdByQuestId(questId)
	qiyuTemp := questtemplate.GetQuestTemplateService().GetQiYuTemplate(qiyuId)
	if qiyuTemp == nil {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	finish := true
	for qiQuestId, _ := range qiyuTemp.GetQuestMap() {
		// 当前任务提交
		if qiQuestId == questId {
			continue
		}

		// 剩余任务是否提交状态
		qu := manager.GetQuestByIdAndState(questtypes.QuestStateCommit, qiQuestId)
		if qu != nil {
			continue
		}

		// 有未提交的任务
		finish = false
	}

	if finish {
		manager.QiYuFinish(qiyuId)
	}

	return nil
}
