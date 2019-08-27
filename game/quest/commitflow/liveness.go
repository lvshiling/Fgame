package commitflow

import (
	livenesslogic "fgame/fgame/game/liveness/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCommitFlow(questtypes.QuestTypeLiveness, quest.CommitFlowHandlerFunc(handleCommitFlowLiveness))
}

//处理活跃度任务
func handleCommitFlowLiveness(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Info("quest:处理活跃度")
	livenesslogic.LivenessQuestCommit(pl, int32(questTemplate.TemplateId()))
	return nil
}
