package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	daliwaneventtypes "fgame/fgame/game/daliwan/event/types"
	daliwanlogic "fgame/fgame/game/daliwan/logic"
	playerdailiwan "fgame/fgame/game/daliwan/player"
	daliwantemplate "fgame/fgame/game/daliwan/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
)

//加载完成后
func daiLiWanUpdate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	obj, ok := data.(*playerdailiwan.DaLiWanObject)
	if !ok {
		return
	}
	linshiTemplate := daliwantemplate.GetDaLiWanTemplateService().GetLinShiTemplate(obj.GetTyp())
	if linshiTemplate == nil {
		return
	}
	daliwanlogic.DaLiWanPropertyChanged(pl)
	scenelogic.AddBuff(pl, linshiTemplate.BuffId, pl.GetId(), common.MAX_RATE)

	return
}

func init() {
	gameevent.AddEventListener(daliwaneventtypes.DaLiWanEventTypeUpdate, event.EventListenerFunc(daiLiWanUpdate))
}
