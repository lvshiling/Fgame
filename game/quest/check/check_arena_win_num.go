package check

import (
	playerarena "fgame/fgame/game/arena/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubType3V3WinNum, quest.CheckHandlerFunc(handleArenaWinNum))
}

//check 处理获得3V3胜利次数
func handleArenaWinNum(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理获得3V3胜利次数")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, _ := range questDemandMap {
		manager := pl.GetPlayerDataManager(types.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
		arenaObj := manager.GetPlayerArenaObject()
		if arenaObj == nil {
			return
		}
		totalNum := arenaObj.GetTotalRewardTime()
		if totalNum == 0 {
			return
		}
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, totalNum)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理获得3V3胜利次数,完成")
	return nil
}
