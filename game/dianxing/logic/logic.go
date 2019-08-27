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
	"fgame/fgame/pkg/mathutils"
)

//变更点星系统属性
func DianXingPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeDianXing.Mask())
	return
}

//点星系统进阶判断
func DianXingAdvanced(pl player.Player, curTimesNum int32, curBless int32, addWfb int32, dianXingTemplate *gametemplate.DianXingTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeDianXing, dianXingTemplate.TimesMin, dianXingTemplate.TimesMax)
	updateRate := dianXingTemplate.UpdateWfb + addWfb
	blessMax := dianXingTemplate.ZhufuMax
	addMin := dianXingTemplate.AddMin
	addMax := dianXingTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)

	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//点星系统解封判断
func DianXingJieFengAdvanced(pl player.Player, curTimesNum int32, curBless int32, jieFengTemplate *gametemplate.DianXingJieFengTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeDianXing, jieFengTemplate.TimesMin, jieFengTemplate.TimesMax)
	updateRate := jieFengTemplate.UpdateWfb
	blessMax := jieFengTemplate.ZhufuMax
	addMin := jieFengTemplate.AddMin
	addMax := jieFengTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)

	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}
