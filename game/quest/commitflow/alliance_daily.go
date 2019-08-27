package commitflow

import (
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCommitFlow(questtypes.QuestTypeDailyAlliance, quest.CommitFlowHandlerFunc(handleCommitFlowDailyAlliance))
}

//处理仙盟日常任务
func handleCommitFlowDailyAlliance(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理仙盟日常任务")

	//随机下一个仙盟日常任务
	questlogic.GetNextDailyQuest(pl, questtypes.QuestDailyTagAlliance)
	return nil
}
