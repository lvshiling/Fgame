package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	xiantaoeventtypes "fgame/fgame/game/xiantao/event/types"
	xiantaologic "fgame/fgame/game/xiantao/logic"
)

//百年仙桃变化
func baiNianXianTaoChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}

	s := p.GetScene()
	if s == nil {
		return
	}

	xiantaologic.PlayerXianTaoChangedBuff(p)
	xiantaologic.PlayerXianTaoInfoChanged(p)
	return
}

func init() {
	gameevent.AddEventListener(xiantaoeventtypes.EventTypeBaiNianXianTaoChange, event.EventListenerFunc(baiNianXianTaoChanged))
}
