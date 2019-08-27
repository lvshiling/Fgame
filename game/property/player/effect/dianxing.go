package effect

import (
	playerdianxing "fgame/fgame/game/dianxing/player"
	dianxingtemplate "fgame/fgame/game/dianxing/template"
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
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeDianXing, DianXingPropertyEffect)
}

//属性作用器点星系统
func DianXingPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeDianXing) {
		return
	}
	dianXingManager := pl.GetPlayerDataManager(types.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	dianXingInfo := dianXingManager.GetDianXingObject()
	percentMap := make(map[playerpropertytypes.PropertyEffectorType]int32)

	dianXingTemplate := dianxingtemplate.GetDianXingTemplateService().GetDianXingTemplateByArg(dianXingInfo.CurrType, dianXingInfo.CurrLevel)
	if dianXingTemplate != nil {
		hp := int64(0)
		attack := int64(0)
		defence := int64(0)

		hp += int64(dianXingTemplate.Hp)
		attack += int64(dianXingTemplate.Attack)
		defence += int64(dianXingTemplate.Defence)

		prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, hp)
		prop.SetBase(propertytypes.BattlePropertyTypeAttack, attack)
		prop.SetBase(propertytypes.BattlePropertyTypeDefend, defence)

		//作用其他模块百分比属性
		percentMap = percentMapValAdd(percentMap, dianXingTemplate.GetExternalPercentMap())
	}

	//解封等级增加属性
	dianXingJieFengTemplate := dianxingtemplate.GetDianXingTemplateService().GetDianXingJieFengTemplateByLev(dianXingInfo.JieFengLev)
	if dianXingJieFengTemplate != nil {
		//基础属性
		oldBaseHp := prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
		oldBaseAttack := prop.GetBase(propertytypes.BattlePropertyTypeAttack)
		oldBaseDefence := prop.GetBase(propertytypes.BattlePropertyTypeDefend)
		prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, int64(dianXingJieFengTemplate.Hp)+oldBaseHp)
		prop.SetBase(propertytypes.BattlePropertyTypeAttack, int64(dianXingJieFengTemplate.Attack)+oldBaseAttack)
		prop.SetBase(propertytypes.BattlePropertyTypeDefend, int64(dianXingJieFengTemplate.Defence)+oldBaseDefence)
		//属性百分比
		oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(dianXingJieFengTemplate.AttrPercent)+oldHp)
		oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(dianXingJieFengTemplate.AttrPercent)+oldAttack)
		oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(dianXingJieFengTemplate.AttrPercent)+oldDefence)
		//作用其他模块百分比属性
		percentMap = percentMapValAdd(percentMap, dianXingJieFengTemplate.GetExternalPercentMap())
	}

	//点星解封套装
	dianXingJieFengTaoZhuangTemplate := dianxingtemplate.GetDianXingTemplateService().GetDianXingJieFengTaoZhuangTemplateByLev(dianXingInfo.JieFengLev)
	if dianXingJieFengTaoZhuangTemplate != nil {
		//基础属性
		oldBaseHp := prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
		oldBaseAttack := prop.GetBase(propertytypes.BattlePropertyTypeAttack)
		oldBaseDefence := prop.GetBase(propertytypes.BattlePropertyTypeDefend)
		prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, int64(dianXingJieFengTaoZhuangTemplate.Hp)+oldBaseHp)
		prop.SetBase(propertytypes.BattlePropertyTypeAttack, int64(dianXingJieFengTaoZhuangTemplate.Attack)+oldBaseAttack)
		prop.SetBase(propertytypes.BattlePropertyTypeDefend, int64(dianXingJieFengTaoZhuangTemplate.Defence)+oldBaseDefence)
		//作用其他模块百分比属性
		percentMap = percentMapValAdd(percentMap, dianXingJieFengTaoZhuangTemplate.GetExternalPercentMap())
	}
	//作用其他模块属性作用器(百分比属性的有作用其他模块为0也需传)
	externalModulePercentEffect(pl, prop, playerpropertytypes.PlayerPropertyEffectorTypeDianXing, percentMap)
}

//作用其他模块属性作用器(百分比属性的有作用其他模块为0也需传)
func externalModulePercentEffect(pl player.Player, prop *propertycommon.SystemPropertySegment, effectorType playerpropertytypes.PropertyEffectorType, percentMap map[playerpropertytypes.PropertyEffectorType]int32) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//先清一下
	for moduleType := range playerpropertytypes.GetPropertyEffectoryTypeMap() {
		module := propertyManager.GetModule(moduleType)
		modulePropertySegment := module.GetExternalPropertySegment(effectorType)
		modulePropertySegment.Clear()
	}

	//获取关联模块
	for moduleType, percent := range percentMap {
		if percent > 0 {
			module := propertyManager.GetModule(moduleType)
			modulePropertySegment := module.GetExternalPropertySegment(effectorType)

			oldHp := modulePropertySegment.Get(uint(propertytypes.BattlePropertyTypeMaxHP))
			modulePropertySegment.Set(uint(propertytypes.BattlePropertyTypeMaxHP), int64(percent)+oldHp)
			oldAttack := modulePropertySegment.Get(uint(propertytypes.BattlePropertyTypeAttack))
			modulePropertySegment.Set(uint(propertytypes.BattlePropertyTypeAttack), int64(percent)+oldAttack)
			oldDefence := modulePropertySegment.Get(uint(propertytypes.BattlePropertyTypeDefend))
			modulePropertySegment.Set(uint(propertytypes.BattlePropertyTypeDefend), int64(percent)+oldDefence)
		}
	}
}

func percentMapValAdd(oldPercentMap map[playerpropertytypes.PropertyEffectorType]int32, addPercentMap map[playerpropertytypes.PropertyEffectorType]int32) map[playerpropertytypes.PropertyEffectorType]int32 {
	for moduleType, percent := range addPercentMap {
		oldPercentMap[moduleType] += percent
	}
	return oldPercentMap
}
