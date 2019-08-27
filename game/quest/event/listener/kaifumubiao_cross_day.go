package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
)

//开服目标过天
func questKaiFuMuBiaoCrossDay(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	questList := questlogic.CheckInitQuestList(pl)
	scQuestListUpdate := pbutil.BuildSCQuestListUpdate(questList)
	pl.SendMsg(scQuestListUpdate)
	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestKaiFuMuBiaoCrossDay, event.EventListenerFunc(questKaiFuMuBiaoCrossDay))
}
