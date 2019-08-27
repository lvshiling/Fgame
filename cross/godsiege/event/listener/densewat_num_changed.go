package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/godsiege/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	gamegodsiegepbutil "fgame/fgame/game/godsiege/pbutil"
	godsiegescene "fgame/fgame/game/godsiege/scene"
	"fgame/fgame/game/scene/scene"
)

//金银密窟采集次数改变
func denseWatNumChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate().(godsiegescene.GodSiegeSceneData)
	godType := sd.GetGodType()
	itemMap := sd.GetItemMapByPlayer(pl)
	scGodSiegeCollectChanged := gamegodsiegepbutil.BuildSCGodSiegeCollectChanged(pl, int32(godType), itemMap)
	pl.SendMsg(scGodSiegeCollectChanged)

	isDenseWatSync := pbutil.BuildISDenseWatSync(pl)
	pl.SendMsg(isDenseWatSync)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerDenseWatNumChanged, event.EventListenerFunc(denseWatNumChanged))
}
