package effect

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playertianshu "fgame/fgame/game/tianshu/player"
	tianshutemplate "fgame/fgame/game/tianshu/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeTianShu, TianShuPropertyEffect)
}

//天书作用器
func TianShuPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	tianshuManager := p.GetPlayerDataManager(types.PlayerTianShuDataManagerType).(*playertianshu.PlayerTianShuDataManager)
	tianshuMap := tianshuManager.GetTianShuAll()
	for typ, obj := range tianshuMap {
		level := obj.GetLevel()
		temp := tianshutemplate.GetTianShuTemplateService().GetTianShuTemplate(typ, level)
		if temp == nil {
			return
		}

		for typ, val := range temp.GetBattleAttrMap() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

}
