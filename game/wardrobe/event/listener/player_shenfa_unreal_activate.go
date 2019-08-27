package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
)

//玩家幻化激活
func playerShenfaUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	shenfaId := data.(int)
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(shenfaId)
	if shenfaTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeShenFa, int32(shenfaId))
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaUnrealActivate, event.EventListenerFunc(playerShenfaUnrealActivate))
}
