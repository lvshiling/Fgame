package listener

import (
	"fgame/fgame/core/event"
	playeractivity "fgame/fgame/game/activity/player"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//退出场景
func battlePlayerExitScene(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	s := p.GetScene()
	if s == nil {
		return
	}

	activityType, flag := s.MapTemplate().GetMapType().ToActivityType()
	if !flag {
		return
	}
	//退出活动
	//保存数据
	activityManager := p.GetPlayerDataManager(types.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityManager.ExitActivity(activityType)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(battlePlayerExitScene))
}
