package check

import (
	playerchess "fgame/fgame/game/chess/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeDragonChess, quest.CheckHandlerFunc(handleQuestFinishDragonChess))
}

//check 完成棋局抽奖X次
func handleQuestFinishDragonChess(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成棋局抽奖X次")

	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	manager := pl.GetPlayerDataManager(types.PlayerChessDataManagerType).(*playerchess.PlayerChessDataManager)
	leftNum := manager.GetAllLeftTimesExcludeGold()
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
		}).Debug("quest:处理完成棋局抽奖X次,完成")
	return nil
}
