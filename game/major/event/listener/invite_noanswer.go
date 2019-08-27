package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	majoreventtypes "fgame/fgame/game/major/event/types"
	"fgame/fgame/game/major/pbutil"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
)

//双修邀请对方无应答
func majorInviteNoAnswer(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}
	majorInvite, ok := data.(*majortypes.MajorInvite)
	if !ok {
		return
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}
	scMajorSpouseRefused := pbutil.BuildSCMajorSpouseRefused(majorInvite.SpouseName, int32(majorInvite.FuBenType), majorInvite.FuBenId)
	pl.SendMsg(scMajorSpouseRefused)
	return
}

func init() {
	gameevent.AddEventListener(majoreventtypes.EventTypeMajorInviteNoAnswer, event.EventListenerFunc(majorInviteNoAnswer))
}
