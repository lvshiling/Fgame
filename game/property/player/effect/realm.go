package effect

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playerrealm "fgame/fgame/game/realm/player"
	realmtemplate "fgame/fgame/game/realm/template"
	skilltypes "fgame/fgame/game/skill/types"
)

// func init() {
// 	playerpropertytypes.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeRealm, RealmPropertyEffect)
// }

// //境界作用器
// func RealmPropertyEffect(p player.Player, prop *propertycommon.BattlePropertySegment) {
// 	//TODO 功能开启

// 	realmManager := p.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
// 	level := realmManager.GetTianJieTaLevel()
// 	tjtTemplate := realmtemplate.GetRealmTemplateService().GetTianJieTaTemplateByLevel(level)
// 	if tjtTemplate == nil {
// 		return
// 	}

// 	//天劫塔属性
// 	for typ, val := range tjtTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
// 		total := prop.Get(typ)
// 		total += val
// 		prop.Set(typ, total)
// 	}

// }

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeRealm, RealmPropertyEffect)
}

//境界作用器
func RealmPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	//TODO 功能开启

	realmManager := p.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	level := realmManager.GetTianJieTaLevel()
	tjtTemplate := realmtemplate.GetRealmTemplateService().GetTianJieTaTemplateByLevel(level)
	if tjtTemplate == nil {
		return
	}

	//天劫塔属性
	for typ, val := range tjtTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
		total := prop.GetBase(typ)
		total += val
		prop.SetBase(typ, total)
	}
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeRealm, prop)
}
