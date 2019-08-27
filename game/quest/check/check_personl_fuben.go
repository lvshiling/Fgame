package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	xianfulogic "fgame/fgame/game/xianfu/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeXianFuPersonal, quest.CheckHandlerFunc(handleQuestFinishXianFuPersonal))
}

//check 完成个人副本
func handleQuestFinishXianFuPersonal(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理个人副本")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, needNum := range questDemandMap {
		leftNum := xianfulogic.TotalFreeTimes(pl)
		if leftNum < needNum {
			if !questTemplate.IsAutoFinishByUsedFree() {
				return
			}

			demandNum := needNum - leftNum
			questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
			flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, demandNum)
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
		}).Debug("quest:处理个人副本,完成")
	return nil
}
