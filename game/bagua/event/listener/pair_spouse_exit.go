package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/bagua/bagua"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//夫妻助战配偶中途退出
func baGuaPairSpouseExit(target event.EventTarget, data event.EventData) (err error) {
	spl, ok := target.(player.Player)
	if !ok {
		return
	}
	bagua.GetBaGuaService().PairSpouseExit(spl.GetId())
	return
}

func init() {
	gameevent.AddEventListener(baguaeventtypes.EventTypeBaGuaPairSpouseExit, event.EventListenerFunc(baGuaPairSpouseExit))
}
