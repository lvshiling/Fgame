package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	winglogic "fgame/fgame/game/wing/logic"
	"fgame/fgame/game/wing/pbutil"
)

//战翼试用过期
func wingTrialOverdue(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*wingeventtypes.WingTrialOverdueEventData)
	if !ok {
		return
	}
	trialOrderId := eventData.GetTrialId()
	bResult := eventData.GetResult()
	winglogic.WingPropertyChanged(pl)
	scWingTrialOverdue := pbutil.BuildSCWingTrialOverdue(trialOrderId, bResult)
	err = pl.SendMsg(scWingTrialOverdue)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingTrialOverdue, event.EventListenerFunc(wingTrialOverdue))
}
