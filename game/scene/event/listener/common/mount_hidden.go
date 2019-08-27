package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	battlelogic "fgame/fgame/game/battle/logic"
	gameevent "fgame/fgame/game/event"
	mountpbutil "fgame/fgame/game/mount/pbutil"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//坐骑隐藏
func mountHidden(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)

	hidden := pl.IsMountHidden()
	scMountHidden := mountpbutil.BuildSCMountHidden(hidden)
	pl.SendMsg(scMountHidden)
	s := pl.GetScene()
	if s == nil {
		return
	}

	scenePlayerMountChanged := pbutil.BuildScenePlayerMountChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerMountChanged)
	battlelogic.UpdateMountBattleProperty(pl)
	lingTong := pl.GetLingTong()
	if lingTong != nil && lingTong.GetScene() == s {
		lingTong.LingTongMountHidden(hidden)
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowMountHidden, event.EventListenerFunc(mountHidden))
}
