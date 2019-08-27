package effect

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeLingyu, LingyuPropertyEffect)
}

//领域作用器
func LingyuPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeLingYu) {
		return
	}
	lingyuManager := p.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()
	advancedId := lingyuInfo.AdvanceId
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(int32(advancedId))
	//领域系统默认不开启 advancedId=0
	if lingyuTemplate == nil {
		return
	}

	//身法属性
	for typ, val := range lingyuTemplate.GetBattlePropertyMap() {
		total := prop.GetBase(typ)
		total += val
		prop.SetBase(typ, total)
	}

	//幻化丹食丹等级
	unrealLevel := lingyuInfo.UnrealLevel
	huanHuaTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuHuanHuaTemplate(unrealLevel)
	if huanHuaTemplate != nil {
		hp := int64(huanHuaTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack := int64(huanHuaTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
		defence := int64(huanHuaTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
	}

	//非进阶领域
	lingYuOtherMap := lingyuManager.GetLingyuOtherMap()
	for _, lingYuTypeOtherMap := range lingYuOtherMap {
		for lingYuOtherId, wo := range lingYuTypeOtherMap {
			//非进阶领域属性
			lingYuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingYuOtherId))
			if lingYuTemplate == nil {
				continue
			}

			skinHp := int64(0)
			skinAttack := int64(0)
			skinDefence := int64(0)
			if lingYuTemplate.FieldUpstarBeginId != 0 && wo.Level != 0 {
				lingYuUpstarTemplate := lingYuTemplate.GetLingYuUpstarByLevel(wo.Level)
				skinHp = int64(lingYuUpstarTemplate.Hp)
				skinAttack = int64(lingYuUpstarTemplate.Attack)
				skinDefence = int64(lingYuUpstarTemplate.Defence)
			}

			for typ, val := range lingYuTemplate.GetBattlePropertyMap() {
				switch typ {
				case propertytypes.BattlePropertyTypeMaxHP:
					{
						val += skinHp
						break
					}
				case propertytypes.BattlePropertyTypeAttack:
					{
						val += skinAttack
						break
					}
				case propertytypes.BattlePropertyTypeDefend:
					{
						val += skinDefence
						break
					}
				}
				total := prop.GetGlobal(typ)
				total += val
				prop.SetGlobal(typ, total)
			}

			if lingYuTemplate.FieldUpstarBeginId != 0 && wo.Level != 0 {
				lingYuUpstarTemplate := lingYuTemplate.GetLingYuUpstarByLevel(wo.Level)

				//领域基础全属性万分比
				if lingYuUpstarTemplate.FieldPercent != 0 {
					oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(lingYuUpstarTemplate.FieldPercent)+oldHp)
					oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(lingYuUpstarTemplate.FieldPercent)+oldAttack)
					oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(lingYuUpstarTemplate.FieldPercent)+oldDefence)
				}
			}

		}
	}
	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeLingYu, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeLingyu, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeLingYuSystemSkill, prop)
}
