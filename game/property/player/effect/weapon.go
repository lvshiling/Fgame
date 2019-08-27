package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playerweapon "fgame/fgame/game/weapon/player"
	"fgame/fgame/game/weapon/weapon"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeWeapon, WeaponPropertyEffect)
}

//兵魂作用器
func WeaponPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeWeapon) {
		return
	}
	weaponManager := p.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	weaponMap := weaponManager.GetAllWeapon()
	for weaponId, weaponInfo := range weaponMap {
		to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
		if to == nil {
			continue
		}
		if weaponInfo.ActiveFlag == 0 {
			continue
		}

		culLevel := weaponInfo.CulLevel
		hp := int64(0)
		attack := int64(0)
		defence := int64(0)
		culTemplate := to.GetWeaponPeiYangByLevel(culLevel)
		if culTemplate != nil {
			hp = int64(culTemplate.Hp)
			attack = int64(culTemplate.Attack)
			defence = int64(culTemplate.Defence)
		}

		if to.GetBattleAttrTemplate() != nil {
			//兵魂属性
			for typ, val := range to.GetBattleAttrTemplate().GetAllBattleProperty() {

				switch typ {
				case propertytypes.BattlePropertyTypeMaxHP:
					{
						val += hp
						break
					}
				case propertytypes.BattlePropertyTypeAttack:
					{
						val += attack
						break
					}
				case propertytypes.BattlePropertyTypeDefend:
					{
						val += defence
						break
					}
				}
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}

		//兵魂升星属性
		if to.WeaponUpgradeBeginId != 0 {
			level := weaponManager.GetWeaponLevel(weaponId)
			if level == 0 {
				continue
			}

			temp := to.GetWeaponUpstarByLevel(level)
			if temp == nil {
				continue
			}

			hp := int64(temp.Hp)
			attack := int64(temp.Attack)
			defence := int64(temp.Defence)

			oldHp := prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
			prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, hp+oldHp)
			oldAtteck := prop.GetBase(propertytypes.BattlePropertyTypeAttack)
			prop.SetBase(propertytypes.BattlePropertyTypeAttack, attack+oldAtteck)
			oldDefence := prop.GetBase(propertytypes.BattlePropertyTypeDefend)
			prop.SetBase(propertytypes.BattlePropertyTypeDefend, defence+oldDefence)
		}

		//觉醒属性
		state := weaponManager.GetWeaponState(weaponId)
		//觉醒
		if state == 1 {
			for typ, val := range to.GetAwakenAttrTemplate().GetAllBattleProperty() {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
	}

}
