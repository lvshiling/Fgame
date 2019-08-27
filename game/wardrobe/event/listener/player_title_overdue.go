package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	titleeventtypes "fgame/fgame/game/title/event/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
)

//玩家活动称号失效
func playerActivityTitleOverdue(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	titleId := data.(int32)
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.RemoveSeqId(wardrobetypes.WardrobeSysTypeTitle, titleId)
	return
}

func init() {
	gameevent.AddEventListener(titleeventtypes.EventTypeTitleActivityOverdue, event.EventListenerFunc(playerActivityTitleOverdue))
}
