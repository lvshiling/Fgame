package effect

import (
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeEquipment, EquipmentBasePropertyEffect)
}

//装备作用器
func EquipmentBasePropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	inventoryManager := p.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equimentBag := inventoryManager.GetEquipmentBag()

	for _, equipment := range equimentBag.GetAll() {
		if equipment.IsEmpty() {
			continue
		}
		//装备属性
		itemTemplate := item.GetItemService().GetItem(int(equipment.ItemId))
		equimentTemplate := itemTemplate.GetEquipmentTemplate()
		//装备自身属性

		for typ, val := range equimentTemplate.GetBattlePropertyMap() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}

	}

	for _, equipment := range equimentBag.GetAll() {
		if equipment.IsEmpty() {
			continue
		}
		//装备属性
		// itemTemplate := item.GetItemService().GetItem(int(equipment.ItemId))

		//强化属性
		level := equipment.Level
		star := equipment.Star
		if star != 0 {
			starTemplate := item.GetItemService().GetEquipStrengthenTemplate(inventorytypes.EquipmentStrengthenTypeStar, equipment.SlotId, star)
			for typ, val := range starTemplate.GetBattlePropertyMap() {
				total := prop.GetGlobal(typ)
				total += val
				prop.SetGlobal(typ, total)
			}
		}
		if level != 0 {
			levelTemplate := item.GetItemService().GetEquipStrengthenTemplate(inventorytypes.EquipmentStrengthenTypeUpgrade, equipment.SlotId, level)
			for typ, val := range levelTemplate.GetBattlePropertyMap() {
				total := prop.GetGlobal(typ)
				total += val
				prop.SetGlobal(typ, total)
			}
		}

		//宝石属性
		for _, itemId := range equipment.GemInfo {
			//宝石属性
			gemItemTemplate := item.GetItemService().GetItem(int(itemId))
			//装备自身属性
			attrTemplate := gemItemTemplate.GetGemAttrTemplate()
			for typ, val := range attrTemplate.GetAllBattleProperty() {
				total := prop.GetGlobal(typ)
				total += val
				prop.SetGlobal(typ, total)
			}
		}

	}

	taoZhuangMap := make(map[*gametemplate.TaozhuangTemplate]int32)
	for _, equipment := range equimentBag.GetAll() {
		if equipment.IsEmpty() {
			continue
		}
		//装备属性
		itemTemplate := item.GetItemService().GetItem(int(equipment.ItemId))
		equimentTemplate := itemTemplate.GetEquipmentTemplate()
		taoZhuangTemplate := equimentTemplate.GetTaozhuangTemplate()
		if taoZhuangTemplate != nil {
			num, exist := taoZhuangMap[taoZhuangTemplate]
			if !exist {
				taoZhuangMap[taoZhuangTemplate] = 1
			} else {
				taoZhuangMap[taoZhuangTemplate] = 1 + num
			}
		}

	}

	//套装属性
	for taoZhuangTemplate, num := range taoZhuangMap {
		if taoZhuangTemplate.Number > num {
			continue
		}
		for typ, val := range taoZhuangTemplate.GetBattlePropertyMap() {
			total := prop.GetGlobal(typ)
			total += val
			prop.SetGlobal(typ, total)
		}
		for typ, val := range taoZhuangTemplate.GetBattlePropertyPercentMap() {
			total := prop.GetGlobalPercent(typ)
			total += val
			prop.SetGlobalPercent(typ, total)
		}
	}

}
