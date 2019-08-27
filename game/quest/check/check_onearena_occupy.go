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
	quest.RegisterCheck(questtypes.QuestSubTypeOneArenaOccupyTime, quest.CheckHandlerFunc(handleOneArenaOccupyTime))
}

//check 处理占领X阶灵池持续Y分钟
func handleOneArenaOccupyTime(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理占领X阶灵池持续Y分钟")
	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, condition := range questDemandMap {
		level := onearenatypes.OneArenaLevelType(demandId)
		if !level.Vaild() {
			continue
		}

		arenaManager := pl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
		if !arenaManager.IsFullQuestCondition(level, int64(condition)) {
			continue
		}
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, condition)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理占领X阶灵池持续Y分钟,完成")
	return nil
}
