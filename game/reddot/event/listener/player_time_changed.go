package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	reddotlogic "fgame/fgame/game/reddot/logic"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
)

//玩家时间条件红点检查
func playerTimeChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	//检查活动红点
	reddotlogic.CheckReddotOnTimeAll(pl)
	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeTimeReddotNotice, event.EventListenerFunc(playerTimeChanged))
}
