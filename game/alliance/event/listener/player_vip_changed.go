package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	vipeventtypes "fgame/fgame/game/vip/event/types"
)

//玩家vip变化
func playerVipChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	if pl.GetAllianceId() == 0 {
		return
	}

	//同步成员信息
	alliance.GetAllianceService().SyncMemberInfo(pl.GetId(), pl.GetName(), pl.GetSex(), pl.GetLevel(), pl.GetForce(), pl.GetZhuanSheng(), pl.GetLingyuInfo().AdvanceId, pl.GetVip())

	return
}

func init() {
	gameevent.AddEventListener(vipeventtypes.EventTypeVipLevelChanged, event.EventListenerFunc(playerVipChanged))
}
