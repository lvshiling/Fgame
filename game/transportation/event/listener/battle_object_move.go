package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	"fgame/fgame/game/transportation/pbutil"
)

//到达目标点
func battleObjectMove(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeBiaoChe {
		return
	}

	biaoChe := n.(*biaochenpc.BiaocheNPC)
	transportationObj := biaoChe.GetTransportationObject()
	//判断是否到达
	if !transportationObj.IsReachGoal(biaoChe.GetPosition()) {
		return
	}

	biaoChe.ReachGoal()

	pl := player.GetOnlinePlayerManager().GetPlayerById(transportationObj.GetPlayerId())
	if pl != nil {
		scTransportBriefInfoNotice := pbutil.BuildSCTransportBriefInfoNotice(transportationObj, n)
		pl.SendMsg(scTransportBriefInfoNotice)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectMove, event.EventListenerFunc(battleObjectMove))
}
