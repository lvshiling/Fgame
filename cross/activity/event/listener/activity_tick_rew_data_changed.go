package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/activity/pbutil"
	"fgame/fgame/game/battle/battle"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//定时奖励数据变化
func battlePlayerActivityTickRewDataChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*battle.BattlePlayerActivityTickRewDataChangedEventData)
	if !ok {
		return
	}

	resMap := eventData.GetAddResMap()
	specialResMap := eventData.GetSpecialResMap()

	isTickRewDataMsg := pbutil.BuildISPlayerActivityTickRewDataChanged(resMap, specialResMap)
	p.SendMsg(isTickRewDataMsg)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityTickRewDataChanged, event.EventListenerFunc(battlePlayerActivityTickRewDataChanged))
}
