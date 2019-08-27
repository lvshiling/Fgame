package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	tianmotemplate "fgame/fgame/game/tianmo/template"
)

//玩家天魔体进阶
func playerTianMoTiAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId := data.(int32)
	template := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(advanceId)
	if template == nil {
		return
	}

	uniteData := tianmoeventtypes.CreatePlayerTianMoTiUnitePiFuEventData(template.GetActivateUniteType(), template.WaiguanValue)
	gameevent.Emit(tianmoeventtypes.EventTypeTianMoUnitePiFu, pl, uniteData)
	return
}

func init() {
	gameevent.AddEventListener(tianmoeventtypes.EventTypeTianMoAdvanced, event.EventListenerFunc(playerTianMoTiAdavanced))
}
