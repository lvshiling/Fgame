package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	"fgame/fgame/game/xianti/xianti"
)

//玩家幻化激活
func playerXianTiUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	xianTiId := data.(int)
	xianTiTemplate := xianti.GetXianTiService().GetXianTi(xianTiId)
	if xianTiTemplate == nil {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeXianTi, int32(xianTiId))
	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiUnrealActivate, event.EventListenerFunc(playerXianTiUnrealActivate))
}
