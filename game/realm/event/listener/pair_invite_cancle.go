package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	"fgame/fgame/game/realm/pbutil"
)

//夫妻助战邀请取消
func pairInviteCancle(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	spouseId := data.(int64)
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl == nil {
		return
	}
	scRealmPairPushCancle := pbutil.BuildSCRealmPairPushCancle(pl.GetName())
	spl.SendMsg(scRealmPairPushCancle)
	return
}

func init() {
	gameevent.AddEventListener(realmeventtypes.EventTypeRealmPairInviteCancle, event.EventListenerFunc(pairInviteCancle))
}
