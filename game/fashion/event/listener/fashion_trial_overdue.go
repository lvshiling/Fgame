package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fashioneventtypes "fgame/fgame/game/fashion/event/types"
	fashionlogic "fgame/fgame/game/fashion/logic"
	"fgame/fgame/game/fashion/pbutil"
	"fgame/fgame/game/player"
)

//玩家时装试用卡过期
func playerFashionTrialOverdue(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*fashioneventtypes.FashionTrialOverdueEventData)
	if !ok {
		return
	}
	fashionId := eventData.GetTrialId()
	overdueType := eventData.GetOverdueType()

	fashionlogic.FashionPropertyChanged(pl)
	scMsg := pbutil.BuildSCFashionTrialOverdueNotice(fashionId, overdueType)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(fashioneventtypes.EventTypeFashionTrialOverdue, event.EventListenerFunc(playerFashionTrialOverdue))
}
