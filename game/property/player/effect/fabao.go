package effect

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotemplate "fgame/fgame/game/fabao/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeFaBao, FaBaoPropertyEffect)
}

//法宝作用器
func FaBaoPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeFaBao) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoInfo := manager.GetFaBaoInfo()
	advancedId := faBaoInfo.GetAdvancedId()
	fabaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(advancedId))
	//法宝系统默认不开启
	if fabaoTemplate == nil {
		return
	}

	for typ, val := range fabaoTemplate.GetBattleProperty() {
		total := prop.GetBase(typ)
		total += val
		prop.SetBase(typ, total)
	}

	//幻化丹食丹等级
	unrealLevel := faBaoInfo.GetUnrealLevel()
	huanHuaTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoHuanHuaTemplate(unrealLevel)
	if huanHuaTemplate != nil {
		hp := int64(huanHuaTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack := int64(huanHuaTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
		defence := int64(huanHuaTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
	}

	//法宝通灵
	tonglingLevel := faBaoInfo.GetTongLingLevel()
	tongLingTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoTongLingTemplate(tonglingLevel)
	if tongLingTemplate != nil {
		oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(tongLingTemplate.FaBaoPercent)+oldHp)
		oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(tongLingTemplate.FaBaoPercent)+oldAttack)
		oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(tongLingTemplate.FaBaoPercent)+oldDefence)
	}

	//非进阶法宝
	faBaoOtherMap := manager.GetFaBaoOtherMap()
	for _, faBaoTypeOtherMap := range faBaoOtherMap {
		for faBaoOtherId, wo := range faBaoTypeOtherMap {
			fabaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoOtherId))

			if fabaoTemplate.FaBaoUpstarBeginId != 0 && wo.GetLevel() != 0 {
				faBaoUpstarTemplate := fabaoTemplate.GetFaBaoUpstarByLevel(wo.GetLevel())
				skinHp := int64(faBaoUpstarTemplate.Hp)
				skinAttack := int64(faBaoUpstarTemplate.Attack)
				skinDefence := int64(faBaoUpstarTemplate.Defence)

				//法宝基础全属性万分比
				if faBaoUpstarTemplate.FaBaoPercent != 0 {
					oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(faBaoUpstarTemplate.FaBaoPercent)+oldHp)
					oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(faBaoUpstarTemplate.FaBaoPercent)+oldAttack)
					oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(faBaoUpstarTemplate.FaBaoPercent)+oldDefence)
				}

				hp := int64(skinHp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
				prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
				attack := int64(skinAttack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
				prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
				defence := int64(skinDefence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
				prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
			}

			for typ, val := range fabaoTemplate.GetBattleProperty() {
				total := prop.GetGlobal(typ)
				total += val
				prop.SetGlobal(typ, total)
			}
		}
	}

	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeFaBao, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeFaBao, prop)
}
