package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/activity/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//活动pk数据变化
func battlePlayerActivityPkDataChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}
	killData := data.(*scene.PlayerActvitiyKillData)
	isPlayerActivityPkDataChanged := pbutil.BuildISPlayerActivityPkDataChanged(killData)
	p.SendMsg(isPlayerActivityPkDataChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityPkDataChanged, event.EventListenerFunc(battlePlayerActivityPkDataChanged))
}
