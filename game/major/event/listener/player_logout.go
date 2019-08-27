package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/major/major"
	"fgame/fgame/game/major/pbutil"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家下线
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	majorInvite := major.GetMajorService().GetMajorInvite(pl.GetId())
	if majorInvite == nil {
		return
	}

	if majorInvite.PlayerId == pl.GetId() {
		spl := player.GetOnlinePlayerManager().GetPlayerById(majorInvite.SpouseId)
		if spl == nil {
			return
		}
		scMajorInvitePushCancle := pbutil.BuildSCMajorInvitePushCancle(pl.GetName(), int32(majorInvite.FuBenType), majorInvite.FuBenId)
		spl.SendMsg(scMajorInvitePushCancle)
	} else {
		spl := player.GetOnlinePlayerManager().GetPlayerById(majorInvite.PlayerId)
		if pl == nil {
			return
		}
		scMajorSpouseRefused := pbutil.BuildSCMajorSpouseRefused(pl.GetName(), int32(majorInvite.FuBenType), majorInvite.FuBenId)
		spl.SendMsg(scMajorSpouseRefused)
	}

	major.GetMajorService().RemoveInvite(majorInvite.PlayerId, majorInvite.SpouseId)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
