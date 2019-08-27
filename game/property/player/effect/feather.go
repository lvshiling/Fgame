package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playerwing "fgame/fgame/game/wing/player"
	"fgame/fgame/game/wing/wing"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeFeather, FeatherPropertyEffect)
}

//护体仙羽作用器
func FeatherPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeFeather) {
		return
	}
	wingManager := p.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	featherId := wingInfo.FeatherId
	//护体仙羽
	featherTemplate := wing.GetWingService().GetFeather(featherId)
	if featherTemplate != nil && featherTemplate.GetBattleAttrTemplate() != nil {
		for typ, val := range featherTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}
}
