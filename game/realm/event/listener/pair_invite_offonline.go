package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	"fgame/fgame/game/realm/pbutil"
)

//夫妻助战闯关途中挑战者掉线
func pairInviteOffonline(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}
	inviteName, ok := data.(string)
	if !ok {
		return
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}
	scRealmInviteOffonline := pbutil.BuildSCRealmInviteOffonline(inviteName)
	pl.SendMsg(scRealmInviteOffonline)
	return
}

func init() {
	gameevent.AddEventListener(realmeventtypes.EventTypeRealmPairInviteOffonline, event.EventListenerFunc(pairInviteOffonline))
}
