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
	quest.RegisterCheck(questtypes.QuestSubTypePlayerLevel, quest.CheckHandlerFunc(handleQuestPlayerLevel))
}

//check 玩家等级达到X级
func handleQuestPlayerLevel(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理玩家等级达到X级")

	level := pl.GetLevel()
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, level)
	if !flag {
		panic("quest:设置 SetQuestData 应该成功")
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理玩家等级达到X级,完成")
	return nil
}
