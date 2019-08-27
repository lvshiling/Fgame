package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
)

//更新副本波数
func playerXianFuRefreshGroup(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	group := data.(int32)
	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubTypeFuBenMonsterGroup, 0, group)
	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuRefreshGroup, event.EventListenerFunc(playerXianFuRefreshGroup))
}
