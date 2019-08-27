package effect

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	skilltypes "fgame/fgame/game/skill/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetemplate "fgame/fgame/game/wardrobe/template"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeWardrobe, WardrobePropertyEffect)
}

//衣橱作用器
func WardrobePropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeWardrobe) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	wardrobeMap := manager.GetWardrobeMap()

	for typ, suitMap := range wardrobeMap {
		mulNum := float64(1)
		peiYangNum := manager.GetWardrobePeiYangNum(typ)
		suitTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSuitTemplate(typ)
		if suitTemplate == nil {
			continue
		}
		peiYangTempalte := suitTemplate.GetPeiYangByLevel(peiYangNum)
		if peiYangTempalte != nil {
			mulNum += (float64(peiYangTempalte.Percent) / float64(common.MAX_RATE))
		}
		for subType, wardrobeObj := range suitMap {
			if !wardrobeObj.GetIsActive() {
				continue
			}
			typ := wardrobeObj.GetType()
			wardrobeTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuTemplate(typ, subType)
			if wardrobeTemplate == nil {
				continue
			}
			for typ, val := range wardrobeTemplate.GetBattlePropertyMap() {
				total := prop.GetBase(typ)
				value := int64(math.Ceil(float64(val) * mulNum))
				total += value
				prop.SetBase(typ, total)
			}
			for typ, val := range wardrobeTemplate.GetBattlePropertyPercentMap() {
				total := prop.GetGlobalPercent(typ)
				value := int64(math.Ceil(float64(val) * mulNum))
				total += value
				prop.SetGlobalPercent(typ, total)
			}
		}
		if peiYangTempalte != nil {
			for typ, val := range peiYangTempalte.GetBattlePropertyMap() {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
	}

	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeWardrobe, prop)
}
