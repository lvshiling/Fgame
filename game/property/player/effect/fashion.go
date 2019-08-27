package effect

import (
	"fgame/fgame/game/fashion/fashion"
	playerfashion "fgame/fgame/game/fashion/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeFashion, FashionPropertyEffect)
}

//时装作用器
func FashionPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeFashion) {
		return
	}
	fashionManager := p.GetPlayerDataManager(types.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	propertyManager := p.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//获取关联模块
	equipmentBaseModule := propertyManager.GetModule(playerpropertytypes.PlayerPropertyEffectorTypeEquipment)
	//获取装备
	fashionPropertySegment := equipmentBaseModule.GetExternalPropertySegment(playerpropertytypes.PlayerPropertyEffectorTypeFashion)
	fashionPropertySegment.Clear()

	for _, fashionTypeMap := range fashionManager.GetFashionMap() {
		for _, fa := range fashionTypeMap {
			fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fa.FashionId))
			//时装系统默认不开启 fashionWear=0
			if fashionTemplate == nil {
				continue
			}
			//时装属性
			if fashionTemplate.GetBattleAttrTemplate() == nil {
				continue
			}

			hp := int64(0)
			attack := int64(0)
			defence := int64(0)
			if fashionTemplate.FashionUpgradeBeginId != 0 && fa.Star != 0 {
				fashionUpstarTemplate := fashionTemplate.GetFashionUpstarByLevel(fa.Star)
				hp = int64(fashionUpstarTemplate.Hp)
				attack = int64(fashionUpstarTemplate.Attack)
				defence = int64(fashionUpstarTemplate.Defence)

				//装备基础全属性万分比
				oldHp := fashionPropertySegment.Get(uint(propertytypes.BattlePropertyTypeMaxHP))
				fashionPropertySegment.Set(uint(propertytypes.BattlePropertyTypeMaxHP), int64(fashionUpstarTemplate.EquipPercent)+oldHp)
				oldAttack := fashionPropertySegment.Get(uint(propertytypes.BattlePropertyTypeAttack))
				fashionPropertySegment.Set(uint(propertytypes.BattlePropertyTypeAttack), int64(fashionUpstarTemplate.EquipPercent)+oldAttack)
				oldDefence := fashionPropertySegment.Get(uint(propertytypes.BattlePropertyTypeDefend))
				fashionPropertySegment.Set(uint(propertytypes.BattlePropertyTypeDefend), int64(fashionUpstarTemplate.EquipPercent)+oldDefence)
			}

			for typ, val := range fashionTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
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
	}

	// 体验卡
	for fashionId, _ := range fashionManager.GetTrialFashionMap() {
		fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
		if fashionTemplate == nil {
			continue
		}
		//时装属性
		if fashionTemplate.GetBattleAttrTemplate() == nil {
			continue
		}

		for typ, val := range fashionTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}
}
