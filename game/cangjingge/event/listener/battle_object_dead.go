package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/cangjingge/pbutil"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//boss死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	//校验怪物类型
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeCangJingGeBoss {
		return
	}

	//boss死亡
	scBroadcast := pbutil.BuildSCCangJingGeBossInfoBroadcast(n)
	for _, pl := range n.GetScene().GetAllPlayers() {
		pl.SendMsg(scBroadcast)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
