package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	"fgame/fgame/game/realm/realm"
)

//夫妻助战配偶中途退出
func realmPairSpouseExit(target event.EventTarget, data event.EventData) (err error) {
	spl, ok := target.(player.Player)
	if !ok {
		return
	}
	realm.GetRealmRankService().PairSpouseExit(spl.GetId())
	return
}

func init() {
	gameevent.AddEventListener(realmeventtypes.EventTypeRealmPairSpouseExit, event.EventListenerFunc(realmPairSpouseExit))
}
