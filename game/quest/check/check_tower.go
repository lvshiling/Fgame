package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	playertower "fgame/fgame/game/tower/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeDaBaoTimePast, quest.CheckHandlerFunc(handleQuestFinishDaBaoTime))
}

//check 完成打宝塔时间
func handleQuestFinishDaBaoTime(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成打宝塔时间")

	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	manager := pl.GetPlayerDataManager(types.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	leftNum := int32(manager.GetRemainTime())
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
		}).Debug("quest:处理完成打宝塔时间,完成")
	return nil
}
