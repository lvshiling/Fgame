package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
)

//玩家幻化激活
func playerLingYuUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fieldId := data.(int)
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(fieldId)
	if lingyuTemplate == nil {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeField, int32(fieldId))
	return
}

func init() {
	gameevent.AddEventListener(lingyueventtypes.EventTypeLingyuUnrealActivate, event.EventListenerFunc(playerLingYuUnrealActivate))
}
