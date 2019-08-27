package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/bagua/bagua"
	"fgame/fgame/game/bagua/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家下线
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	baGuaInvite := bagua.GetBaGuaService().GetBaGuaInvite(pl.GetId())
	if baGuaInvite == nil {
		return
	}

	if baGuaInvite.PlayerId == pl.GetId() {
		spl := player.GetOnlinePlayerManager().GetPlayerById(baGuaInvite.SpouseId)
		if spl == nil {
			return
		}
		scBaGuaPairPushCancle := pbutil.BuildSCBaGuaPairPushCancle(pl.GetName())
		spl.SendMsg(scBaGuaPairPushCancle)
	} else {
		spl := player.GetOnlinePlayerManager().GetPlayerById(baGuaInvite.PlayerId)
		if pl == nil {
			return
		}
		scBaGuaSpouseRefused := pbutil.BuildSCBaGuaSpouseRefused(pl.GetName())
		spl.SendMsg(scBaGuaSpouseRefused)
	}

	bagua.GetBaGuaService().RemoveInvite(baGuaInvite.PlayerId, baGuaInvite.SpouseId)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
