package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	titleeventtypes "fgame/fgame/game/title/event/types"
	"fgame/fgame/game/title/title"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
)

//玩家称号激活
func playerTitleActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	titleId := data.(int32)
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeTitle, titleId)
	return
}

func init() {
	gameevent.AddEventListener(titleeventtypes.EventTypeTitleActivate, event.EventListenerFunc(playerTitleActivate))
}
