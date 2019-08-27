package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/activity/activity"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//场景完成
func sceneFinish(target event.EventTarget, data event.EventData) (err error) {
	s, ok := target.(scene.Scene)
	if !ok {
		return
	}

	activityType, flag := s.MapTemplate().GetMapType().ToActivityType()
	if !flag {
		return
	}

	// 活动结束
	activity.GetActivityService().AddEndRecord(activityType, s.GetEndTime())
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeSceneFinish, event.EventListenerFunc(sceneFinish))
}
