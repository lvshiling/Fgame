package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	playersoul "fgame/fgame/game/soul/player"
	gametemplate "fgame/fgame/game/template"

	clicktypes "fgame/fgame/game/click/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeClickSoulUpgradeButton, quest.CheckHandlerFunc(handleSoulUpgradeButton))
}

//check 处理帝魂点击事件
func handleSoulUpgradeButton(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理帝魂点击事件")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, num := range questDemandMap {
		clickType := clicktypes.ClickSubTypeSoul(demandId)
		if clickType != clicktypes.ClickSubTypeSoulUpgrade {
			return
		}
		manager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
		isAllFull := manager.IfAllUpgradeFull()
		if !isAllFull {
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
		}).Debug("quest:处理帝魂点击事件,完成")
	return nil
}
