package effect

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeSoul, SoulPropertyEffect)
}

//帝魂作用器
func SoulPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	soulManager := p.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	soulInfoMap := soulManager.GetSoulInfoAll()

	//帝魂属性
	for soulTag, soulInfo := range soulInfoMap {
		level := soulInfo.Level
		soulTemplate := soul.GetSoulService().GetSoulTemplateByLevel(soulTag, level)
		strengthenLevel := soulInfo.StrengthenLevel
		strengthenTemplate := soul.GetSoulService().GetSoulStrengthenTemplateByLevel(soulTag, strengthenLevel)

		hp := int64(strengthenTemplate.Hp)
		attack := int64(strengthenTemplate.Attack)
		defence := int64(strengthenTemplate.Defence)
		for typ, val := range soulTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
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

	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeGuHun, prop)
}
