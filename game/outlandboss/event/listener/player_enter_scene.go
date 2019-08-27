package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/outlandboss/pbutil"
	"fgame/fgame/game/outlandboss/outlandboss"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeOutlandBoss {
		return
	}

	//推送场景boss信息
	bossList := outlandboss.GetOutlandBossService().GetOutlandBossListGroupByMap(s.MapId())
	scMsg := pbutil.BuildSCOutlandBossListInfoNotice(bossList)
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
