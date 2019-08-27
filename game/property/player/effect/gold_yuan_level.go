package effect

import (
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeGoldYuan, GoldYuanPropertyEffect)
}

//元神等级作用器
func GoldYuanPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	propertyManager := p.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	goldYuanLevel := propertyManager.GetGoldYuanLevel()

	//元神等级配置
	levelTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetGoldYuanTemplate(goldYuanLevel)
	if levelTemplate == nil {
		return
	}

	for typ, val := range levelTemplate.GetBattlePropertyMap() {
		prop.SetBase(typ, val)
	}
}
