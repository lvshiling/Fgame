package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
)

//防守方虎符改变
func allianceSceneDefendHuFuChanged(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancescene.AllianceSceneData)
	s := sd.GetScene()
	huFu := data.(int64)
	scAllianceSceneDefendHuFuChanged := pbutil.BuildSCAllianceSceneDefendHuFuChanged(huFu)
	s.BroadcastMsg(scAllianceSceneDefendHuFuChanged)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneDefendHuFuChanged, event.EventListenerFunc(allianceSceneDefendHuFuChanged))
}
