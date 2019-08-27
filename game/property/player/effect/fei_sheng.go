package effect

import (
	playerfeisheng "fgame/fgame/game/feisheng/player"
	feishengtemplate "fgame/fgame/game/feisheng/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeFeiShen, feiShengPropertyEffect)
}

//玩家飞升作用器
func feiShengPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiInfo := feiManager.GetFeiShengInfo()

	//飞升等级属性
	feiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(feiInfo.GetFeiLevel())
	for typ, val := range feiTemplate.GetBattleAttrMap() {
		oldVal := prop.GetBase(typ)
		val += oldVal
		prop.SetBase(typ, val)
	}

	//潜能属性
	attrMap := feishengtemplate.GetFeiShengTemplateService().GetQianNengAttrMap(feiInfo.GetTiZhi(), feiInfo.GetLiDao(), feiInfo.GetJinGu())
	for typ, val := range attrMap {
		oldVal := prop.GetBase(typ)
		val += oldVal
		prop.SetBase(typ, val)
	}

}
