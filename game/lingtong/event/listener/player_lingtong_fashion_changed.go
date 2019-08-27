package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	"fgame/fgame/game/player"
)

//灵童时装辩护
func playerLingTongFashionChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	lingTongFashionObj := data.(*playerlingtong.PlayerLingTongFashionObject)

	lingTong := pl.GetLingTong()
	if lingTong == nil {
		return
	}

	fashionId := lingTongFashionObj.GetFashionId()
	lingTong.SetLingTongFashionId(fashionId)

	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeLingTongFashionChanged, event.EventListenerFunc(playerLingTongFashionChanged))
}
