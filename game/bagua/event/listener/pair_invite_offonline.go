package listener

import (
	"fgame/fgame/core/event"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	"fgame/fgame/game/bagua/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
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
	scBaGuaInviteOffonline := pbutil.BuildSCBaGuaInviteOffonline(inviteName)
	pl.SendMsg(scBaGuaInviteOffonline)
	return
}

func init() {
	gameevent.AddEventListener(baguaeventtypes.EventTypeBaGuaPairInviteOffonline, event.EventListenerFunc(pairInviteOffonline))
}
