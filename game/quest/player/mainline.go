package player

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
)

//主线任务
func (pqdm *PlayerQuestDataManager) afterLoadMainLine() (err error) {
	if pqdm.NumOfQuest() == 0 {
		pqdm.initPlayerQuestMap()
	}
	return
}

//第一次初始化QuestMap
func (pqdm *PlayerQuestDataManager) initPlayerQuestMap() (err error) {
	//添加第一个任务
	firstQuestId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBornQuestId)
	pqdm.AddQuest(firstQuestId)

	return
}
