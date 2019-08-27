package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	"fgame/fgame/game/player"
)

//玩家灵童时装过期
func playerLingTongFashionOverdue(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fashionId, ok := data.(int32)
	if !ok {
		return
	}
	lingtonglogic.LingTongFashionPropertyChanged(pl)
	scLingTongFashionRemove := pbutil.BuildSCLingTongFashionRemove(fashionId)
	pl.SendMsg(scLingTongFashionRemove)
	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeLingTongFashionOverdue, event.EventListenerFunc(playerLingTongFashionOverdue))
}
