package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fashioneventtypes "fgame/fgame/game/fashion/event/types"
	"fgame/fgame/game/fashion/fashion"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
)

//玩家时装激活
func playerFashionActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fashionId := data.(int32)
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	if fashionTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeFashion, fashionId)
	return
}

func init() {
	gameevent.AddEventListener(fashioneventtypes.EventTypeFashionActivate, event.EventListenerFunc(playerFashionActivate))
}
