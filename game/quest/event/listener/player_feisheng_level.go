package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	feishengeventtypes "fgame/fgame/game/feisheng/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家飞升等级
func playerFeiShengLevel(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fei, ok := data.(int32)
	if !ok {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeFeiShengLevel, 0, fei)

	// 激活奇遇任务
	questlogic.InitQiYuQuest(pl, fei)
	return
}

func init() {
	gameevent.AddEventListener(feishengeventtypes.EventTypePlayerFeiSheng, event.EventListenerFunc(playerFeiShengLevel))
}
