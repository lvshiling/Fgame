package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	"fgame/fgame/game/quest/pbutil"
	questplayer "fgame/fgame/game/quest/player"
)

//日环任务过5点
func questCrossFiveUpdate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	qu := data.(*questplayer.PlayerQuestObject)
	scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
	pl.SendMsg(scQuestUpdate)
	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestDailyUpdate, event.EventListenerFunc(questCrossFiveUpdate))
}
