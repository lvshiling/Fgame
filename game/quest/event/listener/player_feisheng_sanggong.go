package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	feishengeventtypes "fgame/fgame/game/feisheng/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家苍龙棋局抽奖
func playerFeiShengSanGong(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFeiShengSanGong, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(feishengeventtypes.EventTypePlayerSanGong, event.EventListenerFunc(playerFeiShengSanGong))
}
