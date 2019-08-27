package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	playermassacre "fgame/fgame/game/massacre/player"
	massacretemplate "fgame/fgame/game/massacre/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeMassacre, MassacrePropertyEffect)
}

//戮仙刃作用器
func MassacrePropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeMassacre) {
		return
	}

	massacreManager := p.GetPlayerDataManager(playertypes.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	massacreInfo := massacreManager.GetMassacreInfo()
	advancedId := massacreInfo.AdvanceId
	massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(advancedId)
	//戮仙刃系统默认不开启 advancedId=0
	if massacreTemplate == nil {
		return
	}

	hp := int64(0)
	attack := int64(0)
	defence := int64(0)

	//戮仙刃属性
	hp += int64(massacreTemplate.Hp)
	attack += int64(massacreTemplate.Attack)
	defence += int64(massacreTemplate.Defence)

	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, hp)
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, attack)
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, defence)

}
