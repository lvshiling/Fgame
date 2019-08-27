package effect

import (
	"fgame/fgame/game/bodyshield/bodyshield"
	playerbshield "fgame/fgame/game/bodyshield/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeBodyShield, BodyShieldPropertyEffect)
}

//护体盾作用器
func BodyShieldPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeBodyShield) {
		return
	}

	bodyShieldManager := p.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
	bshieldInfo := bodyShieldManager.GetBodyShiedInfo()
	advancedId := bshieldInfo.AdvanceId
	bodyShieldTemplate := bodyshield.GetBodyShieldService().GetBodyShieldNumber(int32(advancedId))
	//护体盾系统默认不开启 advancedId=0
	if bodyShieldTemplate == nil {
		return
	}
	hp := int64(0)
	attack := int64(0)
	defence := int64(0)

	//护体盾属性
	if bodyShieldTemplate.GetBattleAttrTemplate() != nil {
		for typ, val := range bodyShieldTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
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

	jinJiaDanLevel := bshieldInfo.JinjiadanLevel
	jinJiaDanTemplate := bodyshield.GetBodyShieldService().GetBodyShieldJinJia(jinJiaDanLevel)
	if jinJiaDanTemplate != nil {
		globalHp := int64(jinJiaDanTemplate.Hp)
		globalAttack := int64(jinJiaDanTemplate.Attack)
		globalDefence := int64(jinJiaDanTemplate.Defence)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, globalHp)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, globalAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, globalDefence)
	}

}
