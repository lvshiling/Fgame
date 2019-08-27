package effect

import (
	playerdaliwan "fgame/fgame/game/daliwan/player"
	daliwantemplate "fgame/fgame/game/daliwan/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeDaLiWan, DaLiWanPropertyEffect)
}

//暗器作用器
func DaLiWanPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {

	daLiWanmanager := p.GetPlayerDataManager(playertypes.PlayerDaLiWanDataManagerType).(*playerdaliwan.PlayerDaLiWanManager)

	battleMap := make(map[propertytypes.BattlePropertyType]int64)
	battlePercentMap := make(map[propertytypes.BattlePropertyType]int64)
	now := global.GetGame().GetTimeService().Now()

	for _, obj := range daLiWanmanager.GetDaliWanMap() {
		if obj.IsExpire(now) {
			continue
		}
		linshiTemplate := daliwantemplate.GetDaLiWanTemplateService().GetLinShiTemplate(obj.GetTyp())
		if linshiTemplate == nil {
			continue
		}
		for typ, val := range linshiTemplate.GetBattlePropertyMap() {
			battleMap[typ] += val
		}
		for typ, val := range linshiTemplate.GetBattlePropertyPercentMap() {
			battlePercentMap[typ] += val
		}
	}
	for typ, val := range battleMap {
		prop.SetGlobal(typ, val)
	}
	for typ, val := range battlePercentMap {
		prop.SetGlobalPercent(typ, val)
	}

}
