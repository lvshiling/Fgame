package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	skilltypes "fgame/fgame/game/skill/types"
	playerxuedun "fgame/fgame/game/xuedun/player"
	xueduntemplate "fgame/fgame/game/xuedun/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeXueDun, XueDunPropertyEffect)
}

//血盾作用器
func XueDunPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeXueDu) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	xunDunInfo := manager.GetXueDunInfo()
	culLevel := xunDunInfo.GetCulLevel()
	culTemplate := xueduntemplate.GetXueDunTemplateService().GetXueDunPeiYangTemplate(culLevel)

	if culTemplate == nil {
		return
	}

	for typ, val := range culTemplate.GetBattleProperty() {
		total := prop.GetBase(typ)
		total += val
		prop.SetBase(typ, total)
	}

	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeXueDun, prop)
}
