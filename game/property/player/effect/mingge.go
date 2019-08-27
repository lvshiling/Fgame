package effect

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/item/item"
	playermingge "fgame/fgame/game/mingge/player"
	minggetemplate "fgame/fgame/game/mingge/template"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeMingGe, mingGePropertyEffect)
}

//命格作用器
func mingGePropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeMingGe) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	mingGePanMap := manager.GetMingGePanMap()
	mingGePanRefinedMap := manager.GetMingGePanRefinedMap()
	mingGongMingLiMap := manager.GetMingLiMap()

	//命盘
	for _, mingGeAllSubTypeMap := range mingGePanMap {
		for _, obj := range mingGeAllSubTypeMap {
			mingPanItemMap := obj.GetMingPanItemMap()
			for _, itemId := range mingPanItemMap {
				itemTemplate := item.GetItemService().GetItem(int(itemId))
				if itemTemplate == nil {
					continue
				}
				typeFlag1 := itemTemplate.TypeFlag1
				mingGeTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeTempalte(typeFlag1)
				if mingGeTemplate == nil {
					continue
				}

				for typ, val := range mingGeTemplate.GetBattlePropertyMap() {
					total := prop.GetBase(typ)
					total += val
					prop.SetBase(typ, total)
				}
			}
		}
	}

	//命盘祭炼百分比
	percentPropertyMap := make(map[propertytypes.BattlePropertyType]float64)
	for mingGeType := minggetypes.MingGeTypeNormal; mingGeType <= minggetypes.MingGeTypeSuper; mingGeType++ {
		for mingGeAllSubType, refinedObj := range mingGePanRefinedMap {
			mingGeAllSUbTypeMap := manager.GetMingGePanTypeMap(mingGeType)
			if mingGeAllSUbTypeMap == nil {
				break
			}
			number := refinedObj.GetNumber()
			star := refinedObj.GetStar()
			mingPanTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingPanTemplate(mingGeAllSubType, number, star)
			if mingPanTemplate == nil {
				continue
			}
			percent := int64(mingPanTemplate.Percent)

			//增加percent
			obj, ok := mingGeAllSUbTypeMap[mingGeAllSubType]
			if !ok {
				continue
			}
			mingPanItemMap := obj.GetMingPanItemMap()
			for _, itemId := range mingPanItemMap {
				itemTemplate := item.GetItemService().GetItem(int(itemId))
				if itemTemplate == nil {
					continue
				}
				typeFlag1 := itemTemplate.TypeFlag1
				mingGeTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeTempalte(typeFlag1)
				if mingGeTemplate == nil {
					continue
				}
				for typ, val := range mingGeTemplate.GetBattlePropertyMap() {
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

	//命盘祭炼固定值
	for mingGeAllSubType, refinedObj := range mingGePanRefinedMap {
		number := refinedObj.GetNumber()
		star := refinedObj.GetStar()
		mingPanTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingPanTemplate(mingGeAllSubType, number, star)
		if mingPanTemplate == nil {
			continue
		}

		//固定值
		for typ, val := range mingPanTemplate.GetBattlePropertyMap() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

	//命理
	for mingGongType, _ := range mingGongMingLiMap {
		battlePropertyMap := manager.GetMingLiBattlePropertyMap(mingGongType)
		for typ, val := range battlePropertyMap {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}
}
