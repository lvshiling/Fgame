package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/pbutil"
	"fgame/fgame/game/player"
)

func friendDummyNumChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}

	dummyNum, ok := data.(int32)
	if !ok {
		return
	}

	scMsg := pbutil.BuildSCFriendDummyFriendNumChanged(dummyNum)
	p.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendDummyNumChanged, event.EventListenerFunc(friendDummyNumChanged))
}
