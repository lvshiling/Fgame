package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//玉玺生成
func allianceSceneInitYuXi(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(alliancescene.AllianceSceneData)
	if !ok {
		return
	}
	n, ok := data.(scene.NPC)
	if !ok {
		return
	}

	//
	scMsg := pbutil.BuildSCAllianceSceneYuXiBroadcast(n)
	sd.GetScene().BroadcastMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneInitYuXi, event.EventListenerFunc(allianceSceneInitYuXi))
}
