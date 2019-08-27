package quest_guaji

import (
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

//装备分解
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeGoldEquipmentResolve, guaji.QuestGuaJiFunc(goldEquipmentResolve))
}

//装备分解
func goldEquipmentResolve(pl player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	if len(demandMap) <= 0 {
		return true
	}

	//获取当前任务数据
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	if q == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}
	num := q.QuestDataMap[0]
	totalNum := demandMap[0]
	needNum := totalNum - num
	if needNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务已经完成")
		return true
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	itemIndexList := make([]int32, 0, needNum)
	itemList := inventoryManager.GetBagAll(inventorytypes.BagTypePrim)
	for _, tempItem := range itemList {
		if tempItem.IsEmpty() {
			continue
		}
		itemTemplate := item.GetItemService().GetItem(int(tempItem.ItemId))
		if itemTemplate.IsGoldEquip() {
			itemIndexList = append(itemIndexList, tempItem.Index)
		}
	}
	goldequiplogic.HandleGoldEquipEat(pl, 0, itemIndexList)
	return true
}
