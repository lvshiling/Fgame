package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeFinishTuMo, quest.CheckHandlerFunc(handleQuestFinishTuMo))
}

//check 完成X次屠魔任务
func handleQuestFinishTuMo(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成X次屠魔任务")

	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	leftNum := questManager.GetTuMoLeftDefaultNum()
	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	for demandId, needNum := range questDemandMap {
		if leftNum >= needNum {
			return
		}
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, needNum)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理完成X次屠魔任务,完成")
	return nil
}
