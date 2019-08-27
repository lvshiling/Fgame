package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeSupremeTitle, SupremeTitlePropertyEffect)
}

//至尊称号作用器
func SupremeTitlePropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeDingZhiTitle) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	titleMap := manager.GetTitleMap()
	if len(titleMap) == 0 {
		return
	}

	for titleId, _ := range titleMap {
		titleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(titleId)
		if titleTemplate == nil {
			continue
		}

		for typ, val := range titleTemplate.GetBattleProperty() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}
}
