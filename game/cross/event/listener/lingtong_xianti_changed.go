package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//灵童领域变化
func lingTongXianTiChanged(target event.EventTarget, data event.EventData) (err error) {
	lingTong, ok := target.(scene.LingTong)
	if !ok {
		return
	}
	owner := lingTong.GetOwner()
	pl, ok := owner.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}
	xianTiId := lingTong.GetLingTongXianTiId()
	lingTongXianTiChanged := pbutil.BuildLingTongXianTiChanged(xianTiId)
	pl.SendCrossMsg(lingTongXianTiChanged)

	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeBattleLingTongShowXianTiChanged, event.EventListenerFunc(lingTongXianTiChanged))
}
