package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/activity/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}
	activityType, ok := s.MapTemplate().GetMapType().ToActivityType()
	if ok {
		pl.EnterActivity(activityType, s.GetEndTime())
		//推送场景采集信息
		countMap := pl.GetActivityCollectCountMap(activityType)
		scMsg := pbutil.BuildSCActivityCollectInfoNotice(int32(activityType), countMap)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
