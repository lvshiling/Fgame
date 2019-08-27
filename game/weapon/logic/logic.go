package logic

import (
	commomlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
)

//变更兵魂属性
func WeaponPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeWeapon.Mask())

	return
}

//兵魂培养判断
func WeaponPeiYang(curTimesNum int32, curBless int32, peiYangTemplate *gametemplate.WeaponPeiYangTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := peiYangTemplate.TimesMin
	timesMax := peiYangTemplate.TimesMax
	updateRate := peiYangTemplate.UpdateWfb
	blessMax := peiYangTemplate.ZhufuMax
	addMin := peiYangTemplate.AddMin
	addMax := peiYangTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//兵魂升星判断
func WeaponUpStar(pl player.Player, curTimesNum int32, curBless int32, upStarTemplate *gametemplate.WeaponUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeWeaponUpstar, upStarTemplate.TimesMin, upStarTemplate.TimesMax)
	updateRate := upStarTemplate.UpstarRate
	blessMax := upStarTemplate.ZhufuMax
	addMin := upStarTemplate.AddMin
	addMax := upStarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
