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
	quest.RegisterCheck(questtypes.QuestSubTypeFinishDailyQuestNum, quest.CheckHandlerFunc(handleQuestFinishDailyNum))
}

//check 完成x次日环任务
func handleQuestFinishDailyNum(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成x次日环任务")

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	dailyObj := manager.GetDailyObj(questtypes.QuestDailyTagPerson)
	leftTimes := dailyObj.GetLeftTimes()
	

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	for demandId, num := range questDemandMap {
		if leftTimes < num {
			flag := manager.SetQuestData(int32(questTemplate.TemplateId()), demandId, num)
			if !flag {
				panic("quest:设置 SetQuestData 应该成功")
			}
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理完成x次日环任务,完成")
	return nil
}
