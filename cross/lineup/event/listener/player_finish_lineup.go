package listener

import (
	"fgame/fgame/core/event"
	lineupeventtypes "fgame/fgame/cross/lineup/event/types"
	lineuplogic "fgame/fgame/cross/lineup/logic"
	"fgame/fgame/cross/lineup/pbutil"
	"fgame/fgame/cross/player/player"
	gameevent "fgame/fgame/game/event"
)

//玩家完成排队
func lineupPlayerLineUpFinish(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	eventData := data.(*lineupeventtypes.PlayerLineUpFinishEventData)

	pl := player.GetOnlinePlayerManager().GetPlayerById(eventData.GetPlayerId())
	if pl != nil {
		isMsg := pbutil.BuildISLineupSuccess()
		pl.SendMsg(isMsg)
	}

	lineuplogic.BroadLineUpChanged(-1, lineList, eventData.GetCrossType())
	return
}

func init() {
	gameevent.AddEventListener(lineupeventtypes.EventTypeLineupPlayerLineUpFinish, event.EventListenerFunc(lineupPlayerLineUpFinish))
}
