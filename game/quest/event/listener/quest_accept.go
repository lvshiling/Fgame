package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
)

//任务接受
func questAccept(target event.EventTarget, data event.EventData) (err error) {

	//TODO 添加任务
	questId, ok := data.(int32)
	if !ok {
		return
	}
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	qu, err := questlogic.CheckQuestIfFinish(pl, questId)
	if err != nil {
		return
	}
	if qu != nil {
		scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		pl.SendMsg(scQuestUpdate)
	}

	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestAccept, event.EventListenerFunc(questAccept))
}
