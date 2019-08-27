package effect

import (
	itemservice "fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeWushuangWeapon, WushuangWeaponPropertyEffect)
}

func WushuangWeaponPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	wushuangDataManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	totalHp := int64(0)
	totalAttack := int64(0)
	totalDefence := int64(0)
	for _, slotObj := range wushuangDataManager.GetSlotObjectMap() {
		if !slotObj.IsEquip() {
			continue
		}
		curlevel := slotObj.GetLevel()
		itemId := slotObj.GetItemId()
		itemtemp := itemservice.GetItemService().GetItem(int(itemId))
		if itemtemp == nil {
			continue
		}
		wushuangBaseTemplate := itemtemp.GetWushuangBaseTemplate()
		if wushuangBaseTemplate == nil {
			continue
		}
		wushuangStrengthenTemplate := wushuangBaseTemplate.GetStrengthTemplateByLevel(curlevel)
		if wushuangStrengthenTemplate == nil {
			continue
		}
		totalHp += wushuangBaseTemplate.Hp + wushuangStrengthenTemplate.Hp
		totalAttack += wushuangBaseTemplate.Attack + wushuangStrengthenTemplate.Attack
		totalDefence += wushuangBaseTemplate.Defence + wushuangStrengthenTemplate.Defence
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

	return
}
