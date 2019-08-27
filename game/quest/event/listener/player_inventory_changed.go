package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家物品变更
func playerInventoryChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	//检查可以正在做的任务
	err = questlogic.SetQuestCollectData(pl, questtypes.QuestSubTypeCollectAllItem)
	return
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeInventoryChanged, event.EventListenerFunc(playerInventoryChanged))
}
