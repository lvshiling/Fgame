package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

//世界boss死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	//校验怪物类型
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeWorldBoss {
		return
	}

	//boss死亡
	scWorldBossInfoBroadcast := pbutil.BuildSCWorldBossInfoBroadcast(n, worldbosstypes.BossTypeWorldBoss)
	for _, pl := range n.GetScene().GetAllPlayers() {
		pl.SendMsg(scWorldBossInfoBroadcast)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
