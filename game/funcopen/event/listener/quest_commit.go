package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopenlogic "fgame/fgame/game/funcopen/logic"
	"fgame/fgame/game/funcopen/pbutil"
	"fgame/fgame/game/player"

	propertylogic "fgame/fgame/game/property/logic"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	questeventtypes "fgame/fgame/game/quest/event/types"
)

//任务提交事件
func questCommit(target event.EventTarget, data event.EventData) (err error) {

	p := target.(player.Player)

	updateList, err := funcopenlogic.CheckFuncOpen(p)
	if err != nil {
		return
	}
	if len(updateList) != 0 {
		//TODO 优化
		//更新部分作用器属性

		p.UpdateBattleProperty(playerpropertytypes.PropertyEffectorTypeMaskAll)
		propertylogic.SnapChangedProperty(p)

		scFuncOpenUpdateList := pbutil.BuildSCFuncOpenUpdateList(updateList)
		p.SendMsg(scFuncOpenUpdateList)
	}
	return
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestCommit, event.EventListenerFunc(questCommit))
}
