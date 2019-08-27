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
func lingTongFaBaoChanged(target event.EventTarget, data event.EventData) (err error) {
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
	faBaoId := lingTong.GetLingTongFaBaoId()
	lingTongFaBaoChanged := pbutil.BuildLingTongFaBaoChanged(faBaoId)
	pl.SendCrossMsg(lingTongFaBaoChanged)

	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeBattleLingTongShowFaBaoChanged, event.EventListenerFunc(lingTongFaBaoChanged))
}
