package effect

import (
	"fgame/fgame/game/emperor/emperor"
	emperortemplate "fgame/fgame/game/emperor/template"
	"fgame/fgame/game/player"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"math"
)

func init() {

	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeEmperor, EmperorPropertyEffect)

}

//帝王加成作用器
func EmperorPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	playerId := p.GetId()
	emperorId, robNum := emperor.GetEmperorService().GetEmperorIdAndRobNum()
	if playerId != emperorId {
		return
	}
	weight := emperortemplate.GetEmperorTemplateService().GetEmperorRobCoefficientAttr(robNum)
	dragronChairTemplate := emperortemplate.GetEmperorTemplateService().GetEmperorTemplate()
	firstAttr := dragronChairTemplate.GetFirstAttrTemplate()
	valueAttr := dragronChairTemplate.GetValueAttrTemplate()

	for typ, val := range firstAttr.GetAllBattleProperty() {
		total := prop.GetGlobal(typ)
		total += int64(math.Ceil(float64(val) * weight))
		prop.SetGlobal(typ, total)
	}

	for typ, val := range valueAttr.GetAllBattleProperty() {
		total := prop.GetGlobal(typ)
		total += val
		prop.SetGlobal(typ, total)
	}

}
