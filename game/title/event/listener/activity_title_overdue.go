package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	titleeventtypes "fgame/fgame/game/title/event/types"
	titlelogic "fgame/fgame/game/title/logic"
	"fgame/fgame/game/title/pbutil"
)

//玩家活动称号失效
func playerActivityTitleOverdue(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	activityTitleId := data.(int32)
	titlelogic.TitlePropertyChanged(pl)
	scTitleRemove := pbutil.BuildSCTitleRemove(activityTitleId)
	pl.SendMsg(scTitleRemove)
	return
}

func init() {
	gameevent.AddEventListener(titleeventtypes.EventTypeTitleActivityOverdue, event.EventListenerFunc(playerActivityTitleOverdue))
}
