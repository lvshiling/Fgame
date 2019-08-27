package listener

import (
	"fgame/fgame/core/event"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
)

//玩家噬魂幡进阶
func playerShiHunFanAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId := data.(int32)
	template := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(advanceId)
	if template == nil {
		return
	}

	if template.GetAdvancedUniteType() == commontypes.AdvancedUnitePiFuTypeDefault || template.WaiGuanValue1 == 0 {
		return
	}

	//TODO
	uniteData := shihunfaneventtypes.CreatePlayerShiHunFanUnitePiFuEventData(template.GetAdvancedUniteType(), template.WaiGuanValue1)
	gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanUnitePiFu, pl, uniteData)
	return
}

func init() {
	gameevent.AddEventListener(shihunfaneventtypes.EventTypeShiHunFanAdvanced, event.EventListenerFunc(playerShiHunFanAdavanced))
}
