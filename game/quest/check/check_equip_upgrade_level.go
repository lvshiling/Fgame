package check

import (
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeEquipmentUpgradeLevel, quest.CheckHandlerFunc(handleEquipUpgradeLevel))
}

//check 装备达到X阶
func handleEquipUpgradeLevel(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理装备达到X阶")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for pos, _ := range questDemandMap {
		manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		equipSoltObj := manager.GetEquipByPos(inventorytypes.BodyPositionType(pos))
		if equipSoltObj == nil {
			return
		}
		itemTemplate := item.GetItemService().GetItem(int(equipSoltObj.ItemId))
		if itemTemplate == nil {
			return
		}
		equipTemplate := itemTemplate.GetEquipmentTemplate()
		if equipTemplate == nil {
			return
		}
		level := equipTemplate.Series
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), pos, level)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理装备达到X阶,完成")
	return nil
}
