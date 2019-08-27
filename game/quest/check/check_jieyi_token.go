package check

import (
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeXiongDiToken, quest.CheckHandlerFunc(checkJieYiToken))
}

func checkJieYiToken(pl player.Player, questTemp *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest: 检查玩家结义信物任务是否完成")

	manager := pl.GetPlayerDataManager(types.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if !manager.IsJieYi() {
		return
	}
	level := manager.GetPlayerJieYiObj().GetTokenLevel()

	questDemandMap := questTemp.GetQuestDemandMap(pl.GetRole())
	for demandId, _ := range questDemandMap {
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemp.TemplateId()), demandId, level)
		if !flag {
			panic("quest: 设置 SetQuestData 应该成功")
		}
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest: 检查玩家结义信物任务是否完成,完成")

	return
}
