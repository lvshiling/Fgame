package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cangjingge/cangjingge"
	"fgame/fgame/game/cangjingge/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCangJingGe {
		return
	}

	//推送场景boss信息
	bossList := cangjingge.GetCangJingGeService().GetCangJingGeBossListGroupByMap(s.MapId())
	scMsg := pbutil.BuildSCCangJingGeBossListInfoNotice(bossList)
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
