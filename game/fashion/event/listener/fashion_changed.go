package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fashioneventtypes "fgame/fgame/game/fashion/event/types"
	playerfashion "fgame/fgame/game/fashion/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家时装变化变化
func fashionChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	fashionId := manager.GetFashionId()
	pl.SetFashionId(fashionId)
	return
}

func init() {
	gameevent.AddEventListener(fashioneventtypes.EventTypeFashionChanged, event.EventListenerFunc(fashionChanged))
}
