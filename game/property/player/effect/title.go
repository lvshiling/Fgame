package effect

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playertitle "fgame/fgame/game/title/player"
	"fgame/fgame/game/title/title"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeTitle, TitlePropertyEffect)
}

//称号作用器
func TitlePropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {

	titleManager := p.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)

	for titleId, _ := range titleManager.GetTitleIdMap() {
		titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
		//称号系统默认不开启 titleWear=0
		if titleTemplate == nil {
			continue
		}
		if titleTemplate.GetBattleAttrTemplate() == nil {
			continue
		}

		titleObj := titleManager.GetTitleObjectById(titleId)
		if titleObj == nil {
			continue
		}

		hp := int64(0)
		attack := int64(0)
		defence := int64(0)

		// 称号星级基础属性
		starLev := titleObj.StarLev
		upStarTemp := title.GetTitleService().GetTitleUpStarTemplate(int(titleId), starLev)
		if upStarTemp != nil {
			hp += int64(upStarTemp.Hp)
			attack += int64(upStarTemp.Attack)
			defence += int64(upStarTemp.Defence)
		}

		//称号属性 + 称号星级基础属性
		for typ, val := range titleTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
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

		/*
			// 称号基础全属性万分比
			if upStarTemp.TitlePercent != 0 {
				oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
				prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(upStarTemp.TitlePercent)+oldHp)
				oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
				prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(upStarTemp.TitlePercent)+oldAttack)
				oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
				prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(upStarTemp.TitlePercent)+oldDefence)
			}
		*/

	}

}
