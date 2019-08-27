package quest_guaji

import (
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	gametemplate "fgame/fgame/game/template"

	inventorylogic "fgame/fgame/game/inventory/logic"
	questtypes "fgame/fgame/game/quest/types"

	log "github.com/Sirupsen/logrus"
)

//装备进阶
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeEquipmentUpgradeLevel, guaji.QuestGuaJiFunc(equipmentUpgradeLevel))
}

//装备进阶
func equipmentUpgradeLevel(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {
	guaJiSuccess := false
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))

	for k, v := range demandMap {
		if v <= q.QuestDataMap[k] {
			guaJiSuccess = true
			continue
		}
		pos := inventorytypes.BodyPositionType(k)
		if !pos.Valid() {
			log.WithFields(
				log.Fields{
					"playerId":  p.GetId(),
					"questType": questTemplate.GetQuestSubType().String(),
					"pos":       k,
				}).Warn("quest_guaji:装备进阶,位置无效")
			return false
		}

		flag := equipmentUpgrade(p, pos)
		if !flag {
			continue
		}
		guaJiSuccess = true
	}
	return guaJiSuccess
}

func equipmentUpgrade(p player.Player, pos inventorytypes.BodyPositionType) bool {
	manager := p.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipBag := manager.GetEquipmentBag()
	item := manager.GetEquipByPos(pos)
	//物品不存在
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"pos":      pos.String(),
			}).Warn("quest_guaji:装备进阶,没装备")
		return false
	}

	nextItemTemplate := equipBag.GetNextEquipment(pos)
	if nextItemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"pos":      pos.String(),
			}).Warn("quest_guaji:装备进阶,已经最高阶")
		return false
	}
	nextEquipTemplate := nextItemTemplate.GetEquipmentTemplate()

	//判断消耗条件
	items := nextEquipTemplate.GetNeedItemMap()
	if !manager.HasEnoughItems(items) {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"pos":      pos.String(),
			}).Warn("quest_guaji:装备进阶,物品不足")
		return false
	}
	inventorylogic.EquipSlotUpgrade(p, pos)
	return true
}
