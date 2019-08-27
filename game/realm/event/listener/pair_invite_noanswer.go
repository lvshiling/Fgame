package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	"fgame/fgame/game/realm/pbutil"
)

//夫妻助战邀请对方无应答
func pairInviteNoAnswer(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}
	spouseName, ok := data.(string)
	if !ok {
		return
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}
	scRealmSpouseRefused := pbutil.BuildSCRealmSpouseRefused(spouseName)
	pl.SendMsg(scRealmSpouseRefused)
	return
}

func init() {
	gameevent.AddEventListener(realmeventtypes.EventTypeRealmPairNoAnswer, event.EventListenerFunc(pairInviteNoAnswer))
}
