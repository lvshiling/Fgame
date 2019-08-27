package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/scene/scene"
)

//灵童系统属性变更
func lingTongEnterScene(target event.EventTarget, data event.EventData) (err error) {
	lingTong, ok := target.(scene.LingTong)
	if !ok {
		return
	}
	battleChanged := lingTong.GetSystemBattlePropertyChangedTypesAndReset()
	if len(battleChanged) <= 0 {
		return
	}
	lingTong.Calculate()
	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeBattleLingTongEnterScene, event.EventListenerFunc(lingTongEnterScene))
}
