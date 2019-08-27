package listener

import (
	"fgame/fgame/core/event"
	christmaseventtypes "fgame/fgame/game/christmas/event/types"
	"fgame/fgame/game/christmas/pbutil"
	christmastemplate "fgame/fgame/game/christmas/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//采集物刷新
func christmasRefreshCollect(target event.EventTarget, data event.EventData) (err error) {
	group := data.(int32)

	collectTemp := christmastemplate.GetChristmasTemplateService().GetChristmasTemplate(group)
	if collectTemp == nil {
		return
	}

	s := scene.GetSceneService().GetSceneByMapId(collectTemp.RebornMapId)
	if s == nil {
		return
	}
	scenelogic.RefreshCollectOnScene(s, collectTemp.RebornCount, collectTemp.BiologyId)

	scNoticeMsg := pbutil.BuildSCChristmasCollectRefreshBroadcast(group)
	plList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, player := range plList {
		player.SendMsg(scNoticeMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(christmaseventtypes.EventTypeChristmasRefreshCollect, event.EventListenerFunc(christmasRefreshCollect))
}
