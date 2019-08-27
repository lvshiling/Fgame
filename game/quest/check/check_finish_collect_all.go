package check

import (
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeCollectAllItem, quest.CheckHandlerFunc(handleQuestFinishCollectAllItem))
}

//check b)	收集所有物品
func handleQuestFinishCollectAllItem(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理收集所有物品")

	err = questlogic.SetQuestCollectData(pl, questtypes.QuestSubTypeCollectAllItem)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("quest:处理收集所有物品,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理收集所有物品,完成")
	return nil
}
