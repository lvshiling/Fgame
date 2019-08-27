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
	quest.RegisterCommitFlow(questtypes.QuestTypeDaily, quest.CommitFlowHandlerFunc(handleCommitFlowDaily))
}

//处理日环
func handleCommitFlowDaily(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理日环")

	//随机下一个日环任务
	questlogic.GetNextDailyQuest(pl, questtypes.QuestDailyTagPerson)
	return nil
}
