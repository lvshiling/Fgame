package listener

import (
	"fgame/fgame/core/event"
	playeractivity "fgame/fgame/game/activity/player"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//玩家完成排队
func battlePlayerActivityPkDataSync(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	killData, ok := data.(*scene.PlayerActvitiyKillData)
	if !ok {
		return
	}
	activityManager := p.GetPlayerDataManager(types.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityManager.UpdateActivityPkData(killData.GetActivityType(), killData.GetKilledNum(), killData.GetLastKilledTime())

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityPkDataSync, event.EventListenerFunc(battlePlayerActivityPkDataSync))
}
