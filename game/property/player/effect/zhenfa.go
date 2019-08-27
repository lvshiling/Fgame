package effect

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playerzhenfa "fgame/fgame/game/zhenfa/player"
	zhenfatemplate "fgame/fgame/game/zhenfa/template"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeZhenFa, ZhenFaPropertyEffect)
}

//阵法作用器
func ZhenFaPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeZhenFa) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	zhenFaMap := manager.GetZhenFaMap()
	zhenQiMap := manager.GetZhenQiMap()
	zhenQiXianHuoMap := manager.GetZhenQiXianHuoMap()

	//阵法
	for zhenFaType, obj := range zhenFaMap {
		level := obj.GetLevel()
		zhenFaTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaTempalte(zhenFaType, level)
		if zhenFaTemplate == nil {
			continue
		}
		battlePropertyMap := zhenFaTemplate.GetBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

	//阵旗
	for zhenFaType, zhenFaZhenQiMap := range zhenQiMap {
		for zhenQiType, obj := range zhenFaZhenQiMap {
			number := obj.GetNumber()
			zhenQiTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaZhenQiTemplate(zhenFaType, zhenQiType, number)
			if zhenQiTemplate == nil {
				continue
			}

			battlePropertyMap := zhenQiTemplate.GetBattlePropertyMap()
			for typ, val := range battlePropertyMap {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
	}

	//阵法仙火
	percentPropertyMap := make(map[propertytypes.BattlePropertyType]float64)
	for zhenFaType, obj := range zhenQiXianHuoMap {
		level := obj.GetLevel()
		zhenFaTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaXianHuoTemplate(zhenFaType, level)
		if zhenFaTemplate == nil {
			continue
		}

		battlePropertyMap := zhenFaTemplate.GetBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
		percent := int64(zhenFaTemplate.Percent)
		if zhenFaTemplate.Percent != 0 {
			zhenFaObj := manager.GetZhenFaByType(zhenFaType)
			if zhenFaObj == nil {
				continue
			}
			zhenFaLevel := zhenFaObj.GetLevel()
			var propertyMap map[propertytypes.BattlePropertyType]int64
			if zhenFaLevel == 0 {
				zhenFaJiHuoTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaJiHuoTemplate(zhenFaType)
				if zhenFaJiHuoTemplate == nil {
					continue
				}
				propertyMap = zhenFaJiHuoTemplate.GetBattlePropertyMap()
			} else {
				zhenFaTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaTempalte(zhenFaType, zhenFaLevel)
				if zhenFaTemplate == nil {
					continue
				}
				propertyMap = zhenFaTemplate.GetBattlePropertyMap()
			}
			if len(propertyMap) != 0 {
				for typ, val := range propertyMap {
					percentPropertyMap[typ] += float64(val*percent) / float64(common.MAX_RATE)
				}
			}
		}
	}
	for typ, floatVal := range percentPropertyMap {
		total := prop.GetBase(typ)
		total += int64(math.Ceil(floatVal))
		prop.SetBase(typ, total)
	}

	//阵法套装
	totalLevel := manager.GetAllZhenFaLevel()
	taoZhuangTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaTaoZhuangTemplate(totalLevel)
	if taoZhuangTemplate != nil {
		battlePropertyMap := taoZhuangTemplate.GetBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}
}
