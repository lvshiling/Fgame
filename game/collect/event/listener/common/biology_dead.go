package common

import (
	"fgame/fgame/core/event"
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/mathutils"
)

//采集 死亡打断
func biologyDead(target event.EventTarget, data event.EventData) (err error) {
	n, ok := target.(scene.NPC)
	if !ok {
		return
	}
	parentId := int32(n.GetBiologyTemplate().TemplateId())
	miZangTemplate := n.GetBiologyTemplate().GetMizangTemplate()
	if miZangTemplate == nil {
		return
	}
	flag := mathutils.RandomHit(common.MAX_RATE, int(n.GetBiologyTemplate().BossMizangRate))
	if !flag {
		return
	}
	biologyTemplate := miZangTemplate.GetBiologyTemplate()
	pos := n.GetPosition()
	angle := float64(0)
	miZangNpc := collectnpc.CreateCollectMiZangNPC(scenetypes.OwnerTypeNone, 0, 0, 0, 0, biologyTemplate, pos, angle, miZangTemplate, parentId)
	n.GetScene().AddSceneObject(miZangNpc)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(biologyDead))
}
