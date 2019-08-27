package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//仙盟原地复活次数
func allianceSceneReliveImmediate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	reliveTime, ok := data.(int32)
	if !ok {
		return
	}

	scAllianceSceneReliveTimeChange := pbutil.BuildSCAllianceSceneReliveTimeChange(reliveTime)
	pl.SendMsg(scAllianceSceneReliveTimeChange)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneImmediate, event.EventListenerFunc(allianceSceneReliveImmediate))
}
