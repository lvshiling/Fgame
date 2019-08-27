package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	playersecretcard "fgame/fgame/game/secretcard/player"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeFinishSecretCard, quest.CheckHandlerFunc(handleQuestFinishSecretCard))
}

//check 完成X次天机牌
func handleQuestFinishSecretCard(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成X次天机牌")

	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	leftNum := manager.GetLeftCardNum()
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
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
		}).Debug("quest:处理完成X次天机牌,完成")
	return nil
}
