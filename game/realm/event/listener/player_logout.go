package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/realm/pbutil"
	"fgame/fgame/game/realm/realm"
)

//玩家下线
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	realmInvite := realm.GetRealmRankService().GetRealmInvite(pl.GetId())
	if realmInvite == nil {
		return
	}

	if realmInvite.PlayerId == pl.GetId() {
		spl := player.GetOnlinePlayerManager().GetPlayerById(realmInvite.SpouseId)
		if spl == nil {
			return
		}
		scRealmPairPushCancle := pbutil.BuildSCRealmPairPushCancle(pl.GetName())
		spl.SendMsg(scRealmPairPushCancle)
	} else {
		spl := player.GetOnlinePlayerManager().GetPlayerById(realmInvite.PlayerId)
		if pl == nil {
			return
		}
		scRealmSpouseRefused := pbutil.BuildSCRealmSpouseRefused(pl.GetName())
		spl.SendMsg(scRealmSpouseRefused)
	}

	realm.GetRealmRankService().RemoveInvite(realmInvite.PlayerId, realmInvite.SpouseId)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
