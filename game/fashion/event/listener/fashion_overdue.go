package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fashioneventtypes "fgame/fgame/game/fashion/event/types"
	fashionlogic "fgame/fgame/game/fashion/logic"
	"fgame/fgame/game/fashion/pbutil"
	"fgame/fgame/game/player"
)

//玩家时装过期
func playerFashionOverdue(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fashionId, ok := data.(int32)
	if !ok {
		return
	}
	fashionlogic.FashionPropertyChanged(pl)
	scFashionRemove := pbutil.BuildSCFashionRemove(fashionId)
	pl.SendMsg(scFashionRemove)
	return
}

func init() {
	gameevent.AddEventListener(fashioneventtypes.EventTypeFashionOverdue, event.EventListenerFunc(playerFashionOverdue))
}
