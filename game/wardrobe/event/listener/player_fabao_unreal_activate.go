package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	fabaotemplate "fgame/fgame/game/fabao/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
)

//玩家幻化激活
func playerFaBaoUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	faBaoId := data.(int)
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(faBaoId)
	if faBaoTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeFaBao, int32(faBaoId))
	return
}

func init() {
	gameevent.AddEventListener(fabaoeventtypes.EventTypeFaBaoUnrealActivate, event.EventListenerFunc(playerFaBaoUnrealActivate))
}
