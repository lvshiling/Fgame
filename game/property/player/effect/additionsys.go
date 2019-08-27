package effect

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	propertytypes "fgame/fgame/game/property/types"
)

//属性作用器附加系统属性
func additionSysPropertyEffect(pl player.Player, sysType additionsystypes.AdditionSysType, prop *propertycommon.SystemPropertySegment) {
	// additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	// if !additionsyslogic.GetAdditionSysFuncOpenByType(pl, sysType) {
	// 	return
	// }
	// equipBag := additionsysManager.GetAdditionSysEquipBagByType(sysType)

	// totalHp := int64(0)
	// totalAttack := int64(0)
	// totalDefence := int64(0)
	// for _, slot := range equipBag.GetAll() {
	// 	if slot.IsEmpty() {
	// 		continue
	// 	}
	// 	// slotType := slot.SysType.ConvertToTemplateAdditionSysType()
	// 	//装备属性
	// 	itemId := int(slot.GetItemId())
	// 	equipTemplate := item.GetItemService().GetItem(itemId).GetSystemEquipTemplate()
	// 	totalHp += int64(equipTemplate.GetHp())
	// 	totalAttack += int64(equipTemplate.GetAttack())
	// 	totalDefence += int64(equipTemplate.GetDefence())

	// 	//强化属性
	// 	if slot.Level > 0 {
	// 		strengthenTemplate := additionsystemplate.GetAdditionSysTemplateService().GetBodyStrengthenByArg(slot.SysType, slot.SlotId, slot.Level)
	// 		totalHp += int64(strengthenTemplate.Hp)
	// 		totalAttack += int64(strengthenTemplate.Attack)
	// 		totalDefence += int64(strengthenTemplate.Defence)
	// 	}

	// 	//神铸属性
	// 	if slot.ShenZhuLev > 0 {
	// 		shenZhuTemplate := additionsystemplate.GetAdditionSysTemplateService().GetShenZhuByArg(slot.SlotId, slot.ShenZhuLev)
	// 		totalHp += int64(shenZhuTemplate.Hp)
	// 		totalAttack += int64(shenZhuTemplate.Attack)
	// 		totalDefence += int64(shenZhuTemplate.Defence)
	// 		//神铸万分比属性
	// 		if shenZhuTemplate.Percent > 0 {
	// 			oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
	// 			prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(shenZhuTemplate.Percent)+oldHp)
	// 			oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
	// 			prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(shenZhuTemplate.Percent)+oldAttack)
	// 			oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
	// 			prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(shenZhuTemplate.Percent)+oldDefence)
	// 		}
	// 	}
	// }

	// totalHp += prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
	// totalAttack += prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
	// totalDefence += prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)

	// prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, totalHp)
	// prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, totalAttack)
	// prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, totalDefence)

	// //升级系统属性
	// levelInfo := additionsysManager.GetAdditionSysLevelInfoByType(sysType)
	// if levelInfo.Level > 0 {
	// 	shengJiTemplate := additionsystemplate.GetAdditionSysTemplateService().GetShengJiByArg(sysType, levelInfo.Level)
	// 	if shengJiTemplate != nil {
	// 		g_hp := int64(0)
	// 		g_attack := int64(0)
	// 		g_defence := int64(0)

	// 		g_hp += int64(shengJiTemplate.GetHp()) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
	// 		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, g_hp)
	// 		g_attack += int64(shengJiTemplate.GetAttack()) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
	// 		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, g_attack)
	// 		g_defence += int64(shengJiTemplate.GetDefence()) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
	// 		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, g_defence)
	// 	}
	// }

	// //化灵属性
	// if levelInfo.LingLevel > 0 {
	// 	huaLingTemplate, _ := additionsystemplate.GetAdditionSysTemplateService().GetHuaLingByArg(sysType, levelInfo.LingLevel)
	// 	if huaLingTemplate != nil {
	// 		g_hp := int64(0)
	// 		g_attack := int64(0)
	// 		g_defence := int64(0)

	// 		g_hp += int64(huaLingTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
	// 		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, g_hp)
	// 		g_attack += int64(huaLingTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
	// 		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, g_attack)
	// 		g_defence += int64(huaLingTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
	// 		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, g_defence)
	// 	}
	// }

	// //觉醒属性
	// awakeInfo := additionsysManager.GetAdditionSysAwakeInfoByType(sysType)
	// sysAdvanced := additionsys.GetSystemAdvancedNum(pl, sysType)
	// if awakeInfo.IsAwake > 0 {
	// 	//觉醒加成(1级)
	// 	awakeTemp := additionsystemplate.GetAdditionSysTemplateService().GetAwakeByArg(sysType, sysAdvanced)
	// 	if awakeTemp != nil {
	// 		baseHp := int64(0)
	// 		baseAttack := int64(0)
	// 		baseDefence := int64(0)

	// 		baseHp += int64(awakeTemp.Hp) + prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
	// 		prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, baseHp)
	// 		baseAttack += int64(awakeTemp.Attack) + prop.GetBase(propertytypes.BattlePropertyTypeAttack)
	// 		prop.SetBase(propertytypes.BattlePropertyTypeAttack, baseAttack)
	// 		baseDefence += int64(awakeTemp.Defence) + prop.GetBase(propertytypes.BattlePropertyTypeDefend)
	// 		prop.SetBase(propertytypes.BattlePropertyTypeDefend, baseDefence)
	// 	}
	// 	//觉醒等级加成(大于一级)
	// 	awakeLevelTemp := additionsystemplate.GetAdditionSysTemplateService().GetAwakeLevelByArg(sysType, sysAdvanced, awakeInfo.IsAwake)
	// 	if awakeLevelTemp != nil {
	// 		baseHp := int64(0)
	// 		baseAttack := int64(0)
	// 		baseDefence := int64(0)

	// 		baseHp += int64(awakeLevelTemp.Hp) + prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
	// 		prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, baseHp)
	// 		baseAttack += int64(awakeLevelTemp.Attack) + prop.GetBase(propertytypes.BattlePropertyTypeAttack)
	// 		prop.SetBase(propertytypes.BattlePropertyTypeAttack, baseAttack)
	// 		baseDefence += int64(awakeLevelTemp.Defence) + prop.GetBase(propertytypes.BattlePropertyTypeDefend)
	// 		prop.SetBase(propertytypes.BattlePropertyTypeDefend, baseDefence)
	// 	}
	// }

	// //通灵属性
	// tongLingInfo := additionsysManager.GetAdditionSysTongLingInfoByType(sysType)
	// if tongLingInfo.TongLingLev > 0 {
	// 	//基础全属性万分比
	// 	tongLingTemplate := additionsystemplate.GetAdditionSysTemplateService().GetTongLingByLevel(tongLingInfo.TongLingLev)
	// 	if tongLingTemplate != nil {
	// 		oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
	// 		prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(tongLingTemplate.GetPercent())+oldHp)
	// 		oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
	// 		prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(tongLingTemplate.GetPercent())+oldAttack)
	// 		oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
	// 		prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(tongLingTemplate.GetPercent())+oldDefence)
	// 	}
	// }
	// // taoZhuangSysType := sysType.ConvertToTemplateAdditionSysType()
	// //套装属性
	// taozhuangTemplate := additionsyslogic.GetAdditionSysTaoZhuangByType(pl, sysType)
	// if taozhuangTemplate != nil {
	// 	//坐骑基础全属性万分比
	// 	oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
	// 	prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(taozhuangTemplate.AttrPercent)+oldHp)
	// 	oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
	// 	prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(taozhuangTemplate.AttrPercent)+oldAttack)
	// 	oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
	// 	prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(taozhuangTemplate.AttrPercent)+oldDefence)
	// }

	baseProperty, basePercentProperty, globalProperty, globaPropertyPercent := additionSysPropertyEffectProperty(pl, sysType)
	for typ, val := range baseProperty {
		newVal := prop.GetBase(typ) + val
		prop.SetBase(typ, newVal)
	}
	for typ, val := range basePercentProperty {
		newVal := prop.GetBasePercent(typ) + val
		prop.SetBasePercent(typ, newVal)
	}
	for typ, val := range globalProperty {
		newVal := prop.GetGlobal(typ) + val
		prop.SetGlobal(typ, newVal)
	}
	for typ, val := range globaPropertyPercent {
		newVal := prop.GetGlobalPercent(typ) + val
		prop.SetGlobalPercent(typ, newVal)
	}
}

//属性作用器附加系统属性
func additionSysPropertyEffectProperty(pl player.Player, sysType additionsystypes.AdditionSysType) (baseProperty map[propertytypes.BattlePropertyType]int64, basePercentProperty map[propertytypes.BattlePropertyType]int64, globalProperty map[propertytypes.BattlePropertyType]int64, globalPercentProperty map[propertytypes.BattlePropertyType]int64) {
	baseProperty = make(map[propertytypes.BattlePropertyType]int64)
	basePercentProperty = make(map[propertytypes.BattlePropertyType]int64)
	globalProperty = make(map[propertytypes.BattlePropertyType]int64)
	globalPercentProperty = make(map[propertytypes.BattlePropertyType]int64)

	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	if !additionsyslogic.GetAdditionSysFuncOpenByType(pl, sysType) {
		return
	}
	equipBag := additionsysManager.GetAdditionSysEquipBagByType(sysType)

	totalHp := int64(0)
	totalAttack := int64(0)
	totalDefence := int64(0)
	for _, slot := range equipBag.GetAll() {
		if slot.IsEmpty() {
			continue
		}
		// slotType := slot.SysType.ConvertToTemplateAdditionSysType()
		//装备属性
		itemId := int(slot.GetItemId())
		equipTemplate := item.GetItemService().GetItem(itemId).GetSystemEquipTemplate()
		totalHp += int64(equipTemplate.GetHp())
		totalAttack += int64(equipTemplate.GetAttack())
		totalDefence += int64(equipTemplate.GetDefence())

		//强化属性
		if slot.Level > 0 {
			strengthenTemplate := additionsystemplate.GetAdditionSysTemplateService().GetBodyStrengthenByArg(slot.SysType, slot.SlotId, slot.Level)
			totalHp += int64(strengthenTemplate.Hp)
			totalAttack += int64(strengthenTemplate.Attack)
			totalDefence += int64(strengthenTemplate.Defence)
		}

		//神铸属性
		if slot.ShenZhuLev > 0 {
			shenZhuTemplate := additionsystemplate.GetAdditionSysTemplateService().GetShenZhuByArg(slot.SlotId, slot.ShenZhuLev)
			totalHp += int64(shenZhuTemplate.Hp)
			totalAttack += int64(shenZhuTemplate.Attack)
			totalDefence += int64(shenZhuTemplate.Defence)
			//神铸万分比属性
			if shenZhuTemplate.Percent > 0 {
				basePercentProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(shenZhuTemplate.Percent)
				basePercentProperty[propertytypes.BattlePropertyTypeAttack] += int64(shenZhuTemplate.Percent)
				basePercentProperty[propertytypes.BattlePropertyTypeDefend] += int64(shenZhuTemplate.Percent)

				// oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
				// prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(shenZhuTemplate.Percent)+oldHp)
				// oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
				// prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(shenZhuTemplate.Percent)+oldAttack)
				// oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
				// prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(shenZhuTemplate.Percent)+oldDefence)
			}
		}
	}
	globalProperty[propertytypes.BattlePropertyTypeMaxHP] += totalHp
	globalProperty[propertytypes.BattlePropertyTypeAttack] += totalAttack
	globalProperty[propertytypes.BattlePropertyTypeDefend] += totalDefence

	// totalHp += prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
	// totalAttack += prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
	// totalDefence += prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)

	// prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, totalHp)
	// prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, totalAttack)
	// prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, totalDefence)

	//升级系统属性
	levelInfo := additionsysManager.GetAdditionSysLevelInfoByType(sysType)
	if levelInfo.Level > 0 {
		shengJiTemplate := additionsystemplate.GetAdditionSysTemplateService().GetShengJiByArg(sysType, levelInfo.Level)
		if shengJiTemplate != nil {
			// g_hp := int64(0)
			// g_attack := int64(0)
			// g_defence := int64(0)
			globalProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(shengJiTemplate.GetHp())
			globalProperty[propertytypes.BattlePropertyTypeAttack] += int64(shengJiTemplate.GetAttack())
			globalProperty[propertytypes.BattlePropertyTypeDefend] += int64(shengJiTemplate.GetDefence())

			// g_hp += int64(shengJiTemplate.GetHp()) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
			// prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, g_hp)
			// g_attack += int64(shengJiTemplate.GetAttack()) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
			// prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, g_attack)
			// g_defence += int64(shengJiTemplate.GetDefence()) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
			// prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, g_defence)
		}
	}

	//化灵属性
	if levelInfo.LingLevel > 0 {
		huaLingTemplate, _ := additionsystemplate.GetAdditionSysTemplateService().GetHuaLingByArg(sysType, levelInfo.LingLevel)
		if huaLingTemplate != nil {
			// g_hp := int64(0)
			// g_attack := int64(0)
			// g_defence := int64(0)

			// g_hp += int64(huaLingTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
			// prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, g_hp)
			// g_attack += int64(huaLingTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
			// prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, g_attack)
			// g_defence += int64(huaLingTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
			// prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, g_defence)

			globalProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(huaLingTemplate.Hp)
			globalProperty[propertytypes.BattlePropertyTypeAttack] += int64(huaLingTemplate.Attack)
			globalProperty[propertytypes.BattlePropertyTypeDefend] += int64(huaLingTemplate.Defence)

		}
	}

	//觉醒属性
	awakeInfo := additionsysManager.GetAdditionSysAwakeInfoByType(sysType)
	sysAdvanced := additionsys.GetSystemAdvancedNum(pl, sysType)
	if awakeInfo.IsAwake > 0 {
		//觉醒加成(1级)
		awakeTemp := additionsystemplate.GetAdditionSysTemplateService().GetAwakeByArg(sysType, sysAdvanced)
		if awakeTemp != nil {
			// baseHp := int64(0)
			// baseAttack := int64(0)
			// baseDefence := int64(0)

			// baseHp += int64(awakeTemp.Hp) + prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
			// prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, baseHp)
			// baseAttack += int64(awakeTemp.Attack) + prop.GetBase(propertytypes.BattlePropertyTypeAttack)
			// prop.SetBase(propertytypes.BattlePropertyTypeAttack, baseAttack)
			// baseDefence += int64(awakeTemp.Defence) + prop.GetBase(propertytypes.BattlePropertyTypeDefend)
			// prop.SetBase(propertytypes.BattlePropertyTypeDefend, baseDefence)

			baseProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(awakeTemp.Hp)
			baseProperty[propertytypes.BattlePropertyTypeAttack] += int64(awakeTemp.Attack)
			baseProperty[propertytypes.BattlePropertyTypeDefend] += int64(awakeTemp.Defence)

		}
		//觉醒等级加成(大于一级)
		awakeLevelTemp := additionsystemplate.GetAdditionSysTemplateService().GetAwakeLevelByArg(sysType, sysAdvanced, awakeInfo.IsAwake)
		if awakeLevelTemp != nil {
			// baseHp := int64(0)
			// baseAttack := int64(0)
			// baseDefence := int64(0)
			baseProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(awakeLevelTemp.Hp)
			baseProperty[propertytypes.BattlePropertyTypeAttack] += int64(awakeLevelTemp.Attack)
			baseProperty[propertytypes.BattlePropertyTypeDefend] += int64(awakeLevelTemp.Defence)

			// baseHp += int64(awakeLevelTemp.Hp) + prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
			// prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, baseHp)
			// baseAttack += int64(awakeLevelTemp.Attack) + prop.GetBase(propertytypes.BattlePropertyTypeAttack)
			// prop.SetBase(propertytypes.BattlePropertyTypeAttack, baseAttack)
			// baseDefence += int64(awakeLevelTemp.Defence) + prop.GetBase(propertytypes.BattlePropertyTypeDefend)
			// prop.SetBase(propertytypes.BattlePropertyTypeDefend, baseDefence)
		}
	}

	//通灵属性
	tongLingInfo := additionsysManager.GetAdditionSysTongLingInfoByType(sysType)
	if tongLingInfo.TongLingLev > 0 {
		//基础全属性万分比
		tongLingTemplate := additionsystemplate.GetAdditionSysTemplateService().GetTongLingByLevel(tongLingInfo.TongLingLev)
		if tongLingTemplate != nil {
			basePercentProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(tongLingTemplate.GetPercent())
			basePercentProperty[propertytypes.BattlePropertyTypeAttack] += int64(tongLingTemplate.GetPercent())
			basePercentProperty[propertytypes.BattlePropertyTypeDefend] += int64(tongLingTemplate.GetPercent())

			// oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
			// prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(tongLingTemplate.GetPercent())+oldHp)
			// oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
			// prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(tongLingTemplate.GetPercent())+oldAttack)
			// oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
			// prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(tongLingTemplate.GetPercent())+oldDefence)
		}
	}
	// taoZhuangSysType := sysType.ConvertToTemplateAdditionSysType()
	//套装属性
	taozhuangTemplate := additionsyslogic.GetAdditionSysTaoZhuangByType(pl, sysType)
	if taozhuangTemplate != nil {
		//坐骑基础全属性万分比
		// oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
		// prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(taozhuangTemplate.AttrPercent)+oldHp)
		// oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
		// prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(taozhuangTemplate.AttrPercent)+oldAttack)
		// oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
		// prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(taozhuangTemplate.AttrPercent)+oldDefence)

		basePercentProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(taozhuangTemplate.AttrPercent)
		basePercentProperty[propertytypes.BattlePropertyTypeAttack] += int64(taozhuangTemplate.AttrPercent)
		basePercentProperty[propertytypes.BattlePropertyTypeDefend] += int64(taozhuangTemplate.AttrPercent)

	}
	return
}
