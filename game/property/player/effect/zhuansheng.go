package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeZhuanSheng, ZhuanShengPropertyEffect)
}

//转生作用器
func ZhuanShengPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeYuanGodGold) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curZhuanSheng := manager.GetZhuanSheng()
	zhuanShengTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetZhuanShengTemplate(curZhuanSheng)
	if zhuanShengTemplate == nil {
		return
	}
	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, int64(zhuanShengTemplate.Hp))
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, int64(zhuanShengTemplate.Attack))
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, int64(zhuanShengTemplate.Defence))
}
