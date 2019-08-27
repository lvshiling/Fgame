package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家城战获胜
func playerAllianceWinChengZhan(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	winNum := data.(int32)
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeChengZhanWinNum, 0, winNum)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceWinChengZhan, event.EventListenerFunc(playerAllianceWinChengZhan))
}
