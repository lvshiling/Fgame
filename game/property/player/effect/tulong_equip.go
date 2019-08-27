package effect

import (
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	playertulongequip "fgame/fgame/game/tulongequip/player"
	tulongequiptemplate "fgame/fgame/game/tulongequip/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeTuLongEquip, tuLongEquipPropertyEffect)
}

//屠龙装备属性
func tuLongEquipPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	allSlotMap := tulongequipManager.GetAllEquipSlotMap()

	totalHp := int64(0)
	totalAttack := int64(0)
	totalDefence := int64(0)

	for suitType, slotList := range allSlotMap {
		for _, slot := range slotList {
			if slot.IsEmpty() {
				continue
			}
			itemId := int(slot.GetItemId())
			level := slot.GetLevel()
			pos := slot.GetSlotId()
			tulongEquipTemp := item.GetItemService().GetItem(itemId).GetTuLongEquipTemplate()
			if tulongEquipTemp == nil {
				continue
			}

			//装备属性
			totalHp += int64(tulongEquipTemp.Hp)
			totalAttack += int64(tulongEquipTemp.Attack)
			totalDefence += int64(tulongEquipTemp.Defence)

			//强化属性
			strengthenTemp := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipStrengthenTemplate(suitType, pos, level)
			if strengthenTemp != nil {
				totalHp += int64(strengthenTemp.Hp)
				totalAttack += int64(strengthenTemp.Attack)
				totalDefence += int64(strengthenTemp.Defence)
			}
		}
	}

	curHp := prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
	curHp += totalHp
	curAttack := prop.GetBase(propertytypes.BattlePropertyTypeAttack)
	curAttack += totalAttack
	curDefend := prop.GetBase(propertytypes.BattlePropertyTypeDefend)
	curDefend += totalDefence
	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, curHp)
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, curAttack)
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, curDefend)

	skillByModulePropertyEffect(pl, skilltypes.SkillFirstTypeTuLongEquip, prop)
	skillByModulePropertyEffect(pl, skilltypes.SkillFirstTypeTuLongEquipSuit, prop)
}
