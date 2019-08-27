package check

import (
	playeronearena "fgame/fgame/game/onearena/player"
	onearenatypes "fgame/fgame/game/onearena/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubType1V1, quest.CheckHandlerFunc(handleQuestOneArena))
}

//check 参加灵池争夺x次
func handleQuestOneArena(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理参加灵池争夺x次")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, num := range questDemandMap {
		manager := pl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
		oneArenaObj := manager.GetOneArena()
		oneArenaLevel := oneArenaObj.Level
		if oneArenaLevel < onearenatypes.OneArenaLevelMax {
			return
		}

		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, num)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理参加灵池争夺x次,完成")
	return nil
}
