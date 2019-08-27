package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	"fgame/fgame/game/godsiege/pbutil"
	godsiegescene "fgame/fgame/game/godsiege/scene"
	"fgame/fgame/game/scene/scene"
)

func denseWatItemChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	sd := data.(godsiegescene.GodSiegeSceneData)

	godType := sd.GetGodType()
	itemMap := sd.GetItemMapByPlayer(pl)

	scGodSiegeCollectChanged := pbutil.BuildSCGodSiegeCollectChanged(pl, int32(godType), itemMap)
	pl.SendMsg(scGodSiegeCollectChanged)

	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeDenseWatItemChanged, event.EventListenerFunc(denseWatItemChanged))
}
