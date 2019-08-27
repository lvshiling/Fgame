package logic

import (
	commomlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	gametemplate "fgame/fgame/game/template"
)

//变更阵法属性
func ZhenFaPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeZhenFa.Mask())
	return
}

//阵法升级
func ZhenFaShengJi(curTimesNum int32, curBless int32, zhenFaTemplate *gametemplate.ZhenFaTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := zhenFaTemplate.TimesMin
	timesMax := zhenFaTemplate.TimesMax
	updateRate := zhenFaTemplate.UpdateWfb
	blessMax := zhenFaTemplate.ZhuFuMax
	addMin := zhenFaTemplate.AddMin
	addMax := zhenFaTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//阵旗进阶
func ZhenQiAdvanced(curTimesNum int32, curBless int32, zhenQiTemplate *gametemplate.ZhenFaZhenQiTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := zhenQiTemplate.TimesMin
	timesMax := zhenQiTemplate.TimesMax
	updateRate := zhenQiTemplate.UpdateWfb
	blessMax := zhenQiTemplate.ZhuFuMax
	addMin := zhenQiTemplate.AddMin
	addMax := zhenQiTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//阵法仙火升级
func ZhenFaXianHuoShengJi(curTimesNum int32, curBless int32, zhenFaTemplate *gametemplate.ZhenFaXianHuoTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := zhenFaTemplate.TimesMin
	timesMax := zhenFaTemplate.TimesMax
	updateRate := zhenFaTemplate.UpdateWfb
	blessMax := zhenFaTemplate.ZhuFuMax
	addMin := zhenFaTemplate.AddMin
	addMax := zhenFaTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
