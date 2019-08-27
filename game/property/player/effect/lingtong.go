package effect

import (
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeLingTong, LingTongPropertyEffect)
}

//灵童作用器
func LingTongPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeLingTong) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongMap := manager.GetLingTongMap()

	for lingTongId, lingTongObj := range lingTongMap {
		//单个灵童属性自己计算
		singleLingTong(p, lingTongId, lingTongObj, prop)
	}

	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeLingTongNormal, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeLingTongSkill, prop)
}

func singleLingTong(p player.Player, lingTongId int32, lingTongObj *playerlingtong.PlayerLingTongInfoObject, prop *propertycommon.SystemPropertySegment) {
	lingtongBaseProperty := make(map[propertytypes.BattlePropertyType]int64)
	lingtongBasePercentProperty := make(map[propertytypes.BattlePropertyType]int64)
	lingtongGlobalProperty := make(map[propertytypes.BattlePropertyType]int64)
	lingtongGlobalPercentProperty := make(map[propertytypes.BattlePropertyType]int64)
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	battlePropertyMap := lingTongTemplate.GetBattleProperty()
	lingtongBaseProperty = mergePropertyMap(lingtongBaseProperty, battlePropertyMap)
	// for typ, val := range battlePropertyMap {
	// 	total := prop.GetBase(typ)
	// 	total += val
	// 	prop.SetBase(typ, total)
	// }

	//培养丹食丹等级
	culLevel := lingTongObj.GetPeiYangLevel()
	lingTongPeiYangTemplate := lingTongTemplate.GetLingTongPeiYangByLevel(culLevel)
	if lingTongPeiYangTemplate != nil {
		battlePropertyMap := lingTongPeiYangTemplate.GetBattleProperty()
		lingtongGlobalProperty = mergePropertyMap(lingtongGlobalProperty, battlePropertyMap)
		// for typ, val := range battlePropertyMap {
		// 	value := val + prop.GetGlobal(typ)
		// 	prop.SetGlobal(typ, value)
		// }
	}

	//升级
	level := lingTongObj.GetLevel()
	lingTongShengJiTemplate := lingTongTemplate.GetLingTongShengJiByLevel(level)
	if lingTongShengJiTemplate != nil {
		battlePropertyMap := lingTongShengJiTemplate.GetBattleProperty()
		lingtongBaseProperty = mergePropertyMap(lingtongBaseProperty, battlePropertyMap)
		// for typ, val := range battlePropertyMap {
		// 	value := val + prop.GetBase(typ)
		// 	prop.SetBase(typ, value)
		// }
	}

	//升星
	starLevel := lingTongObj.GetStarLevel()
	lingTongUpstarTemplate := lingTongTemplate.GetLingTongUpstarByLevel(starLevel)
	if lingTongUpstarTemplate != nil {
		battlePropertyMap := lingTongUpstarTemplate.GetBattlePropertyMap()
		lingtongGlobalProperty = mergePropertyMap(lingtongGlobalProperty, battlePropertyMap)
		// for typ, val := range battlePropertyMap {
		// 	value := val + prop.GetGlobal(typ)
		// 	prop.SetGlobal(typ, value)
		// }
	}

	sysType, ok := additionsystypes.ConvertLingTongIdToAdditionSysType(int(lingTongId))
	if ok {

		//灵珠
		// totalHp := prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
		// totalAttack := prop.GetBase(propertytypes.BattlePropertyTypeAttack)
		// totalDefend := prop.GetBase(propertytypes.BattlePropertyTypeDefend)
		// totalHpPercent := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
		// totalAttackPercent := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
		// totalDefencePercent := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
		totalHp := int64(0)
		totalAttack := int64(0)
		totalDefend := int64(0)
		totalHpPercent := int64(0)
		totalAttackPercent := int64(0)
		totalDefencePercent := int64(0)
		additionsysManager := p.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
		lingZhuMap := additionsysManager.GetAdditionSysLingZhuMap(sysType)
		for _, lingZhuObj := range lingZhuMap {
			lingzhuTemplate := additionsystemplate.GetAdditionSysTemplateService().GetLingZhuTemplate(lingZhuObj.GetLingZhuType())
			if lingzhuTemplate == nil {
				continue
			}
			lingzhuLevelTemplate := lingzhuTemplate.GetLevelTemplate(lingZhuObj.GetLevel())
			if lingzhuLevelTemplate == nil {
				continue
			}
			totalHp += int64(lingzhuLevelTemplate.Hp)
			totalAttack += int64(lingzhuLevelTemplate.Attack)
			totalDefend += int64(lingzhuLevelTemplate.Defence)
			totalHpPercent += int64(lingzhuLevelTemplate.Percent)
			totalAttackPercent += int64(lingzhuLevelTemplate.Percent)
			totalDefencePercent += int64(lingzhuLevelTemplate.Percent)
		}
		//全局属性
		lingtongGlobalProperty[propertytypes.BattlePropertyTypeMaxHP] += totalHp
		lingtongGlobalProperty[propertytypes.BattlePropertyTypeAttack] += totalAttack
		lingtongGlobalProperty[propertytypes.BattlePropertyTypeDefend] += totalDefend
		//基础属性百分比
		lingtongBasePercentProperty[propertytypes.BattlePropertyTypeMaxHP] += totalHpPercent
		lingtongBasePercentProperty[propertytypes.BattlePropertyTypeAttack] += totalAttackPercent
		lingtongBasePercentProperty[propertytypes.BattlePropertyTypeDefend] += totalDefencePercent
		// prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, totalHp)
		// prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, totalAttack)
		// prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, totalDefend)
		// prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, totalHpPercent)
		// prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, totalAttackPercent)
		// prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, totalDefencePercent)

		//灵童装备
		baseProperty, basePercentProperty, globalProperty, globaPropertyPercent := additionSysPropertyEffectProperty(p, sysType)
		lingtongBaseProperty = mergePropertyMap(lingtongBaseProperty, baseProperty)
		lingtongBasePercentProperty = mergePropertyMap(lingtongBasePercentProperty, basePercentProperty)
		lingtongGlobalProperty = mergePropertyMap(lingtongGlobalProperty, globalProperty)
		lingtongGlobalPercentProperty = mergePropertyMap(lingtongGlobalPercentProperty, globaPropertyPercent)

	}

	for typ, val := range lingtongBaseProperty {
		percent := lingtongBasePercentProperty[typ]
		newVal := prop.GetBase(typ) + int64(math.Ceil(float64(val)*float64(int64(common.MAX_RATE)+percent)/float64(common.MAX_RATE)))
		prop.SetBase(typ, newVal)
	}
	for typ, val := range lingtongGlobalProperty {
		percent := lingtongGlobalPercentProperty[typ]
		newVal := prop.GetGlobal(typ) + int64(math.Ceil(float64(val)*float64(int64(common.MAX_RATE)+percent)/float64(common.MAX_RATE)))
		prop.SetGlobal(typ, newVal)
	}
}

func mergePropertyMap(parent, child map[propertytypes.BattlePropertyType]int64) map[propertytypes.BattlePropertyType]int64 {
	for key, val := range child {
		_, ok := parent[key]
		if !ok {
			parent[key] = val
		} else {
			parent[key] += val
		}
	}
	return parent
}
