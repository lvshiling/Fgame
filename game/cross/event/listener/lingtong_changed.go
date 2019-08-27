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
func lingTongChanged(target event.EventTarget, data event.EventData) (err error) {
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
	lingTongId := lingTong.GetLingTongId()
	lingTongName := lingTong.GetName()
	lingTongChanged := pbutil.BuildLingTongChanged(lingTongId, lingTongName)
	pl.SendCrossMsg(lingTongChanged)

	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeBattleLingTongChanged, event.EventListenerFunc(lingTongChanged))
}
