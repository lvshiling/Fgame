package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	xueduneventtypes "fgame/fgame/game/xuedun/event/types"
	"fgame/fgame/game/xuedun/pbutil"
)

//玩家血盾血炼值改变
func playerXueDunBloodChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	xueDunBlood := data.(int64)
	scXueDunBlood := pbutil.BuildSCXueDunBlood(xueDunBlood)
	p.SendMsg(scXueDunBlood)
	return
}

func init() {
	gameevent.AddEventListener(xueduneventtypes.EventTypeXueDunBloodChanged, event.EventListenerFunc(playerXueDunBloodChanged))
}
