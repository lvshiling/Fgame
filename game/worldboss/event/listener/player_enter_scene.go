package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeWorldBoss {
		return
	}

	//推送场景boss信息
	bossList := worldboss.GetWorldBossService().GetWorldBossListGroupByMap(s.MapId())
	scWorldBossListInfoNotice := pbutil.BuildSCWorldBossListInfoNotice(bossList, worldbosstypes.BossTypeWorldBoss, 0)
	pl.SendMsg(scWorldBossListInfoNotice)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
