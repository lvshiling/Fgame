package logic

import (
	commomlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	gametemplate "fgame/fgame/game/template"
)

//变更血盾属性
func XueDunPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeXueDun.Mask())
	return
}

//血盾升阶判断
func XueDunUpgrade(curTimesNum int32, curBless int32, upgradeTemplate *gametemplate.BloodShieldTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upgradeTemplate.TimesMin
	timesMax := upgradeTemplate.TimesMax
	updateRate := upgradeTemplate.UpdatePercent
	blessMax := upgradeTemplate.ZhufuMax
	addMin := upgradeTemplate.AddMin
	addMax := upgradeTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
