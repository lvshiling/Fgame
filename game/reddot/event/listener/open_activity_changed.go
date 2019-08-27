package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	reddotlogic "fgame/fgame/game/reddot/logic"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	playerwelfare "fgame/fgame/game/welfare/player"
)

// 玩家活动数据变化
func openActivityChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	obj, ok := data.(*playerwelfare.PlayerOpenActivityObject)
	if !ok {
		return
	}

	//检查活动红点
	reddotlogic.ReddotNoticeChanged(pl, []*playerwelfare.PlayerOpenActivityObject{obj})
	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeActivityDataChanged, event.EventListenerFunc(openActivityChanged))
}
