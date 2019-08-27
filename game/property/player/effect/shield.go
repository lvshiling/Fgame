package effect

import (
	"fgame/fgame/game/bodyshield/bodyshield"
	playerbshield "fgame/fgame/game/bodyshield/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeShield, ShieldPropertyEffect)
}

//神盾尖刺作用器
func ShieldPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeShield) {
		return
	}
	bodyShieldManager := p.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
	bshieldInfo := bodyShieldManager.GetBodyShiedInfo()
	shieldId := bshieldInfo.ShieldId

	//神盾尖刺
	shieldTemplate := bodyshield.GetBodyShieldService().GetShield(shieldId)
	if shieldTemplate != nil && shieldTemplate.GetBattleAttrTemplate() != nil {
		for typ, val := range shieldTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

}
