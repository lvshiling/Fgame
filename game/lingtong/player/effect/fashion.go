package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongplayerproperty "fgame/fgame/game/lingtong/player/property"
	lingtongplayertypes "fgame/fgame/game/lingtong/player/types"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
)

func init() {
	lingtongplayerproperty.RegisterLingTongPropertyEffector(lingtongplayertypes.LingTongPropertyEffectorTypeFashion, FashionPropertyEffect)
}

//初始作用器
func FashionPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeLingTongHuanHua) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	activateMap := manager.GetActivateFashionMap()
	fashionTrialObj := manager.GetFashionTrialObject()

	if fashionTrialObj != nil && !fashionTrialObj.GetIsExpire() {
		fashionId := fashionTrialObj.GetTrialFashionId()
		lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
		if lingTongFashionTemplate != nil {
			battlePropertyMap := lingTongFashionTemplate.GetLingTongBattlePropertyMap()
			for typ, val := range battlePropertyMap {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
	}

	for _, fashionObj := range activateMap {
		fashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionObj.GetFashionId())
		for typ, val := range fashionTemplate.GetLingTongBattlePropertyMap() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

}
