package listener

import (
	"fgame/fgame/core/event"
	bufftemplate "fgame/fgame/game/buff/template"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家跨服进入
func crossEnter(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	for _, b := range pl.GetBuffs() {
		buffId := b.GetBuffId()
		buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
		if buffTemplate.GetBuffType() == scenetypes.BuffTypeTitleDingZhi {
			continue
		}
		if buffTemplate.GetOfflineSaveType() == scenetypes.BuffOfflineSaveTypeNone {
			pl.RemoveBuff(buffId)
		}
	}
	return nil
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossEnter, event.EventListenerFunc(crossEnter))
}
