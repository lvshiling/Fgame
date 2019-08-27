package effect

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeAnqi, AnqiPropertyEffect)
}

//暗器作用器
func AnqiPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeAnQi) {
		return
	}

	anqiManager := p.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiInfo := anqiManager.GetAnqiInfo()
	advancedId := anqiInfo.AdvanceId
	anqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(advancedId))
	//暗器系统默认不开启 advancedId=0
	if anqiTemplate == nil {
		return
	}
	hp := int64(0)
	attack := int64(0)
	defence := int64(0)

	//暗器属性
	hp += int64(anqiTemplate.Hp)
	attack += int64(anqiTemplate.Attack)
	defence += int64(anqiTemplate.Defence)

	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, hp)
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, attack)
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, defence)

	//暗器食丹等级
	anqiDanLevel := anqiInfo.AnqiDanLevel
	anqiDanTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiDan(anqiDanLevel)
	if anqiDanTemplate != nil {
		globalHp := int64(anqiDanTemplate.Hp)
		globalAttack := int64(anqiDanTemplate.Attack)
		globalDefence := int64(anqiDanTemplate.Defence)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, globalHp)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, globalAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, globalDefence)
	}

	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeAnqiJiguan, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeAnQi, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeAnQiJiGuan, prop)
}
