package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	xuechipbutil "fgame/fgame/game/xuechi/pbutil"
)

//从血池补血
func xueChiRecover(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	recover, ok := data.(int64)
	if !ok {
		return
	}

	pl.AddHP(recover)
	scObjectDamage := pbutil.BuildSCObjectDamage(pl, scenetypes.DamageTypeXueChi, recover, 0, 0)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scObjectDamage)

	blood := pl.GetBlood()
	scXueChiBlood := xuechipbutil.BuildSCXueChiBlood(blood)
	pl.SendMsg(scXueChiBlood)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerXueChiRecover, event.EventListenerFunc(xueChiRecover))
}
