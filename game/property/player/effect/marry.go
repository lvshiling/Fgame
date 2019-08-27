package effect

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	"math"
)

// func init() {

// 	playerpropertytypes.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeMarry, MarryPropertyEffect)
// }

// //作用器
// func MarryPropertyEffect(p player.Player, prop *propertycommon.BattlePropertySegment) {
// 	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeMarry) {
// 		return
// 	}
// 	marryManager := p.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
// 	marryInfo := marryManager.GetMarryInfo()
// 	if marryInfo.Status == marrytypes.MarryStatusTypeUnmarried {
// 		return
// 	}

// 	ringType := marryInfo.Ring
// 	ringLevel := marryInfo.RingLevel
// 	treeLevel := marryInfo.TreeLevel

// 	ringTemplate := marrytemplate.GetMarryTemplateService().GetMarryRingTemplate(ringType, ringLevel)
// 	if ringTemplate.GetBattleAttrTemplate() != nil {
// 		for typ, val := range ringTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
// 			total := prop.Get(typ)
// 			total += val
// 			prop.Set(typ, total)
// 		}
// 	}

// 	treeTemplate := marrytemplate.GetMarryTemplateService().GetMarryLoveTreeTemplate(treeLevel)
// 	if treeTemplate.GetBattleAttrTemplate() != nil {
// 		for typ, val := range treeTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
// 			total := prop.Get(typ)
// 			total += val
// 			prop.Set(typ, total)
// 		}
// 	}
// }

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeMarry, MarryPropertyEffect)
}

//作用器
func MarryPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeMarry) {
		return
	}
	marryManager := p.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := marryManager.GetMarryInfo()

	// 表白系统
	if marryInfo.Status == marrytypes.MarryStatusTypeMarried {
		developLevel := marryManager.GetMarryDevelopLevel()
		developTemplate := marrytemplate.GetMarryTemplateService().GetMarryDeveopTemplate(developLevel)
		if developTemplate != nil {
			for typ, val := range developTemplate.GetBattleProperty() {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
		//配偶提供
		coupleDevelLevel := marryManager.GetCoupleMarryDevelopLevel()
		coupleDevelopTemplate := marrytemplate.GetMarryTemplateService().GetMarryDeveopTemplate(coupleDevelLevel)
		if coupleDevelopTemplate != nil {
			for typ, val := range coupleDevelopTemplate.GetBattleProperty() {
				total := prop.GetBase(typ)
				total += int64(math.Ceil(float64(val) * float64(coupleDevelopTemplate.Percent) / float64(common.MAX_RATE)))
				prop.SetBase(typ, total)
			}
		}
	}

	//增加定情信物套装属性
	addDingQingPropertyEffect(p, prop)

	if marryInfo.Status != marrytypes.MarryStatusTypeMarried {
		return
	}

	ringType := marryInfo.Ring
	ringLevel := marryInfo.RingLevel
	treeLevel := marryInfo.TreeLevel

	ringTemplate := marrytemplate.GetMarryTemplateService().GetMarryRingTemplate(ringType, ringLevel)
	if ringTemplate.GetBattleAttrTemplate() != nil {
		for typ, val := range ringTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

	treeTemplate := marrytemplate.GetMarryTemplateService().GetMarryLoveTreeTemplate(treeLevel)
	if treeTemplate != nil {
		if treeTemplate.GetBattleAttrTemplate() != nil {
			for typ, val := range treeTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
	}

}

//结婚定情
func addDingQingPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeYingLingPu) {
		return
	}
	hp := int64(0)
	attack := int64(0)
	defence := int64(0)

	marryManager := p.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)

	playerSuitMap := marryManager.GetAllDingQingMap()
	if len(playerSuitMap) == 0 {
		return
	}
	spouseSuitMap := marryManager.GetSpouseSuit()

	for suitId, posMap := range playerSuitMap {
		//己算碎片
		for posId, _ := range posMap {
			item := marrytemplate.GetMarryTemplateService().GetMarryXinWuItem(suitId, posId)
			if item == nil {
				continue
			}
			addTimes := int32(1)
			_, exists := spouseSuitMap[suitId]
			if exists {
				_, exists = spouseSuitMap[suitId][posId]
				if exists {
					addTimes = int32(2)
				}
			}
			hp += int64(item.Hp * addTimes)
			attack += int64(item.Attack * addTimes)
			defence += int64(item.Defence * addTimes)
		}

		//开始计算套装
		suitLen := len(posMap)
		spouseLen := 0
		_, exists := spouseSuitMap[suitId]
		if exists {
			spouseLen = len(spouseSuitMap[suitId])
		}
		suitTemplate := marrytemplate.GetMarryTemplateService().GetMarryXinWuGroupTemplate(suitId)
		if suitTemplate == nil {
			continue
		}
		for i := 1; i <= suitLen; i++ {
			suitAddMap := suitTemplate.GetSuitAddMap()
			_, exists := suitAddMap[int32(i)]
			if exists {
				suitItem := suitAddMap[int32(i)]
				suitTime := int32(1)
				if i <= spouseLen { //伴侣也有
					suitTime = int32(2)
				}
				hp += int64(suitItem.Hp * suitTime)
				attack += int64(suitItem.Attack * suitTime)
				defence += int64(suitItem.Defence * suitTime)
			}
		}
	}
	hp += prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
	attack += prop.GetBase(propertytypes.BattlePropertyTypeAttack)
	defence += prop.GetBase(propertytypes.BattlePropertyTypeDefend)

	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, hp)
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, attack)
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, defence)
}
