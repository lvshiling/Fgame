package logic

import (
	commomlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	gametemplate "fgame/fgame/game/template"
)

//变更灵童时装属性属性
func LingTongFashionPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeLingTongFashion.Mask())
	//更新灵童自身属性
	LingTongSelfFashionPropertyChanged(pl)
	return
}

//灵童时装升星判断
func LingTongFashionUpstar(curTimesNum int32, curBless int32, fashionStarTemplate *gametemplate.LingTongFashionUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := fashionStarTemplate.TimesMin
	timesMax := fashionStarTemplate.TimesMax
	updateRate := fashionStarTemplate.UpdateWfb
	blessMax := fashionStarTemplate.ZhufuMax
	addMin := fashionStarTemplate.AddMin
	addMax := fashionStarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
