package listener

import (
	"fgame/fgame/core/event"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	"fgame/fgame/game/arenapvp/pbutil"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//积分变化推送
func arenapvpJiFenChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	arenapvpObj, ok := data.(*playerarenapvp.PlayerArenapvpObject)
	if !ok {
		return
	}

	jiFen := arenapvpObj.GetJiFen()
	scMsg := pbutil.BuildSCArenapvpJiFenChanged(jiFen)
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpJiFenChanged, event.EventListenerFunc(arenapvpJiFenChanged))
}
