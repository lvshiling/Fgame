package check

import (
	playermaterial "fgame/fgame/game/material/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypechallengeMaterialFuBen, quest.CheckHandlerFunc(handleChallengeMaterialFuBen))
}

//check 材料副本
func handleChallengeMaterialFuBen(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理材料副本")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for damandId, needNum := range questDemandMap {
		manager := pl.GetPlayerDataManager(types.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
		leftNum := manager.GetAllLeftTimes()
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		if leftNum < needNum {
			if !questTemplate.IsAutoFinishByUsedFree() {
				return
			}
			flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), damandId, needNum)
			if !flag {
				panic("quest:设置 SetQuestData 应该成功")
			}
			return
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理材料副本,完成")
	return nil
}
