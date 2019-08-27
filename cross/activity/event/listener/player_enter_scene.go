package listener

import (
	"fgame/fgame/core/event"
	activitytemplate "fgame/fgame/game/activity/template"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()

	activityType, ok := s.MapTemplate().GetMapType().ToActivityType()
	if ok {
		now := global.GetGame().GetTimeService().Now()
		activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activityType)
		timeTemplate, _ := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
		endTime, _ := timeTemplate.GetEndTime(now) 

		pl.EnterActivity(activityType, endTime)
	}

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
