package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/player"

	marryeventtypes "fgame/fgame/game/marry/event/types"
)

//玩家婚戒替换
func playerRingFeedUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	ringLevel := data.(int32)
	marry.GetMarryService().MarryRingLevel(pl.GetId(), ringLevel)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryRingFeedUpgrade, event.EventListenerFunc(playerRingFeedUpgrade))
}
