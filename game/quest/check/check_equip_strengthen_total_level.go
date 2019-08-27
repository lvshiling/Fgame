package check

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeEquipmentStrengthenTotalLevel, quest.CheckHandlerFunc(handleEquipStrengthenTotalLevel))
}

//check 装备强化总等级
func handleEquipStrengthenTotalLevel(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理装备强化总等级")

	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	totalLevel := manager.GetEquipTotalLevel()
	if totalLevel == 0 {
		return
	}
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, totalLevel)
	if !flag {
		panic("quest:设置 SetQuestData 应该成功")
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理装备强化总等级,完成")
	return nil
}
