package listener

import (
	"fgame/fgame/core/event"
	christmaseventtypes "fgame/fgame/game/christmas/event/types"
	christmastemplate "fgame/fgame/game/christmas/template"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//采集物清除
func christmasStopCollect(target event.EventTarget, data event.EventData) (err error) {
	group := data.(int32)

	collectTemp := christmastemplate.GetChristmasTemplateService().GetChristmasTemplate(group)
	if collectTemp == nil {
		return
	}

	s := scene.GetSceneService().GetSceneByMapId(collectTemp.RebornMapId)
	if s == nil {
		return
	}
	scenelogic.ClearCollectOnScene(s, collectTemp.BiologyId)
	return
}

func init() {
	gameevent.AddEventListener(christmaseventtypes.EventTypeChristmasStopCollect, event.EventListenerFunc(christmasStopCollect))
}
