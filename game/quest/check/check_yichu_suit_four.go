package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	playerwardrobe "fgame/fgame/game/wardrobe/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeActivateYiChuSuitFour, quest.CheckHandlerFunc(handleActivateYiChuSuitFour))
}

//check 处理激活4条属性的衣橱套装X件
func handleActivateYiChuSuitFour(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理激活4条属性的衣橱套装X件")
	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, _ := range questDemandMap {
		manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
		activateNum := manager.GetWardrobeActivateTypeNumByNum(3)
		if activateNum == 0 {
			return
		}
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, activateNum)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理激活1条属性的衣橱套装X件,完成")
	return nil
}
