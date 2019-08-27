package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

// 活动期间玩家累积消费元宝
func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerGoldCostIncludeBind, event.EventListenerFunc(playerCostGold))
}

func playerCostGold(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	goldNum, ok := data.(int64)
	if !ok {
		return
	}

	if goldNum <= 0 {
		return
	}

	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeCostGold, 0, int32(goldNum))
	return
}
