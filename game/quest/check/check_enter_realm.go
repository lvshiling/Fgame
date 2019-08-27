package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	playerrealm "fgame/fgame/game/realm/player"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeEnterRealm, quest.CheckHandlerFunc(handleEnterRealm))
}

//check 玩家进入天劫塔
func handleEnterRealm(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理玩家进入天劫塔")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, num := range questDemandMap {
		manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
		flag := manager.IfFullLevel()
		if !flag {
			return
		}
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag = questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, num)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理玩家进入天劫塔,完成")
	return nil
}
