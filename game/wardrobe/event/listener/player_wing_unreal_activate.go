package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	"fgame/fgame/game/wing/wing"
)

//玩家幻化激活
func playerWingUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	wingId := data.(int)
	wingTemplate := wing.GetWingService().GetWing(wingId)
	if wingTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeWing, int32(wingId))
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingUnrealActivate, event.EventListenerFunc(playerWingUnrealActivate))
}
