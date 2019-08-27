package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	"fgame/fgame/game/player"
)

//玩家时装试用卡过期
func playerLingTongFashionTrialOverdue(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*lingtongeventtypes.LingTongFashionTrialOverdueEventData)
	if !ok {
		return
	}
	fashionId := eventData.GetTrialId()
	overdueType := eventData.GetOverdueType()
	lingtonglogic.LingTongFashionPropertyChanged(pl)

	scMsg := pbutil.BuildSCLingTongFashionTrialOverdueNotice(fashionId, int32(overdueType))
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeLingTongFashionTrialOverdue, event.EventListenerFunc(playerLingTongFashionTrialOverdue))
}
