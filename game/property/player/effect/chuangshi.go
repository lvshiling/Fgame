package effect

// import (
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	chuangshitemplate "fgame/fgame/game/chuangshi/template"
// 	"fgame/fgame/game/player"
// 	"fgame/fgame/game/player/types"
// 	propertycommon "fgame/fgame/game/property/common"
// 	playerpropertyproperty "fgame/fgame/game/property/player/property"
// 	playerpropertytypes "fgame/fgame/game/property/player/types"
// )

// func init() {
// 	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeChuangShi, ChuangShiPropertyEffect)
// }

// //创世之战作用器
// func ChuangShiPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
// 	chuangShiManager := p.GetPlayerDataManager(types.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	guanZhiInfo := chuangShiManager.GetPlayerChuangShiGuanZhiInfo()

// 	level := guanZhiInfo.GetLevel()
// 	guanZhiTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiGuanZhiTemplate(level)
// 	if guanZhiTemp != nil {
// 		for typ, val := range guanZhiTemp.GetBattlePropertyMap() {
// 			total := prop.GetBase(typ)
// 			total += val
// 			prop.SetBase(typ, total)
// 		}
// 	}
// }
