package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	xueduneventtypes "fgame/fgame/game/xuedun/event/types"
	playerxuedun "fgame/fgame/game/xuedun/player"
)

//玩家血盾阶数改变
func playerXueDunNumberChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	xueDunInfo := manager.GetXueDunInfo()
	number := xueDunInfo.GetNumber()
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeXueDun), number)
	return
}

func init() {
	gameevent.AddEventListener(xueduneventtypes.EventTypeXueDunNumberChanged, event.EventListenerFunc(playerXueDunNumberChanged))
}
