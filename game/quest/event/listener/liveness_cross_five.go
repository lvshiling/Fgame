package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
)

//任务过5点
func questLivenessCrossFive(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	//活跃度任务
	questList := questlogic.CheckInitQuestList(pl)
	scQuestListUpdate := pbutil.BuildSCQuestListUpdate(questList)
	pl.SendMsg(scQuestListUpdate)
	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestLivenessCrossFive, event.EventListenerFunc(questLivenessCrossFive))
}
