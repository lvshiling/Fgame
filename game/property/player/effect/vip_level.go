package effect

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playervip "fgame/fgame/game/vip/player"
	viptemplate "fgame/fgame/game/vip/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeVipLevel, VipLevelPropertyEffect)
}

//VIP等级作用器
func VipLevelPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	vipManager := p.GetPlayerDataManager(types.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)

	//vip等级配置
	level, star := vipManager.GetVipLevel()
	temp := viptemplate.GetVipTemplateService().GetVipTemplate(level, star)
	if temp == nil {
		return
	}

	for typ, val := range temp.GetBattleAttrMap() {
		curVal := prop.GetBase(typ)
		val += curVal
		prop.SetBase(typ, val)
	}
}
