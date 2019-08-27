package listener

import (
	"fgame/fgame/core/event"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	"fgame/fgame/game/bagua/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
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
	scBaGuaPairPushCancle := pbutil.BuildSCBaGuaPairPushCancle(pl.GetName())
	spl.SendMsg(scBaGuaPairPushCancle)
	return
}

func init() {
	gameevent.AddEventListener(baguaeventtypes.EventTypeBaGuaPairInviteCancle, event.EventListenerFunc(pairInviteCancle))
}
