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
	quest.RegisterCommitFlow(questtypes.QuestTypeKaiFuMuBiao, quest.CommitFlowHandlerFunc(handleCommitFlowKaiFuMuBiao))
}

//处理开服目标
func handleCommitFlowKaiFuMuBiao(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理开服目标")
	if questTemplate.GetQuestType() != questtypes.QuestTypeKaiFuMuBiao {
		return
	}
	questId := int32(questTemplate.TemplateId())
	kaiFuDayList, flag := questtemplate.GetQuestTemplateService().GetKaiFuMuBiaoKaiFuDay(questId)
	if !flag {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	manager.AddFinishNum(kaiFuDayList)
	return nil
}
