package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	majoreventtypes "fgame/fgame/game/major/event/types"
	"fgame/fgame/game/major/pbutil"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
)

//双修邀请取消
func majorInviteCancle(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	majorInvite, ok := data.(*majortypes.MajorInvite)
	if !ok {
		return
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(majorInvite.SpouseId)
	if spl == nil {
		return
	}
	scMajorInvitePushCancle := pbutil.BuildSCMajorInvitePushCancle(pl.GetName(), int32(majorInvite.FuBenType), majorInvite.FuBenId)
	spl.SendMsg(scMajorInvitePushCancle)
	return
}

func init() {
	gameevent.AddEventListener(majoreventtypes.EventTypeMajorInviteCancle, event.EventListenerFunc(majorInviteCancle))
}
