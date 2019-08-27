package listener

import (
	"fgame/fgame/core/event"
	buffcommon "fgame/fgame/game/buff/common"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

func buffUpdate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}
	buffObject := data.(buffcommon.BuffObject)
	//定制称号
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffObject.GetBuffId())
	if buffTemplate == nil {
		return
	}

	siBuffUpdate := pbutil.BuildSIBuffUpdate(buffObject)
	pl.SendCrossMsg(siBuffUpdate)

	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffUpdate, event.EventListenerFunc(buffUpdate))
}
