package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/dingshi/dingshi"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

//boss死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	s := bo.GetScene()
	if s == nil {
		return
	}

	//校验怪物类型
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeDingShiBoss {
		return
	}
	//记录死亡时间
	dingshi.GetDingShiService().BossDead(s.MapId(), int32(n.GetBiologyTemplate().TemplateId()))
	//boss死亡
	scBroadcast := pbutil.BuildSCWorldBossInfoBroadcast(n, worldbosstypes.BossTypeDingShi)
	for _, pl := range n.GetScene().GetAllPlayers() {
		pl.SendMsg(scBroadcast)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
