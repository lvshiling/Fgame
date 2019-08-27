package quest_guaji

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	"math"

	log "github.com/Sirupsen/logrus"
)

//装备强化
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeEquipmentStrengthenLevel, guaji.QuestGuaJiFunc(equipmentStrengthenLevel))
}

//装备强化
func equipmentStrengthenLevel(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

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
				}).Warn("quest_guaji:装备强化位置无效")
			return false
		}

		flag := equipmentStrengthenLevelPos(p, pos)
		if !flag {
			continue
		}
		guaJiSuccess = true
	}
	return guaJiSuccess
}

func equipmentStrengthenLevelPos(p player.Player, pos inventorytypes.BodyPositionType) bool {

	manager := p.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	item := manager.GetEquipByPos(pos)
	//物品不存在
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"pos":      pos.String(),
			}).Warn("quest_guaji:装备强化,没装备")
		return false
	}

	propertyManager := p.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	equipmentBag := manager.GetEquipmentBag()

	//判断槽位是否可以升星
	nextEquipmentStrengthenTemplate := equipmentBag.GetNextUpgradeEquipStrengthenTemplate(pos)
	if nextEquipmentStrengthenTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"pos":      pos.String(),
			}).Warn("quest_guaji:装备强化,已经满级")
		return false
	}
	maxLevel := int32(math.Floor(float64(p.GetLevel()) / float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEquipmentStrengthenLevelLimit))))
	//判断等级限制
	if nextEquipmentStrengthenTemplate.Level > maxLevel {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"pos":      pos.String(),
			}).Warn("quest_guaji:装备强化,达到极限")

		return false
	}

	items := nextEquipmentStrengthenTemplate.GetNeedItemMap()
	if len(items) != 0 {
		if !manager.HasEnoughItems(items) {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"pos":      pos.String(),
				}).Warn("quest_guaji:装备强化,物品不足")
			return false
		}
	}
	//判断是否有足够的银两
	if nextEquipmentStrengthenTemplate.SilverNum > 0 {
		if !propertyManager.HasEnoughSilver(int64(nextEquipmentStrengthenTemplate.SilverNum)) {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"pos":      pos.String(),
				}).Warn("quest_guaji:装备强化,银两不足")
			return false
		}
	}

	_, flag := inventorylogic.EquipmentSlotStrengthenUpgrade(p, pos, false)
	if !flag {
		return false
	}
	return true
}
