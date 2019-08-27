package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	"fgame/fgame/game/transportation/pbutil"
	transportationplayer "fgame/fgame/game/transportation/player"
	transportationtypes "fgame/fgame/game/transportation/types"
)

//镖车被攻击
func transportationBeAttack(target event.EventTarget, data event.EventData) (err error) {

	defenceObj := target.(scene.BattleObject)
	attackObj := data.(scene.BattleObject)
	n, ok := defenceObj.(scene.NPC)
	if !ok {
		return
	}
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeBiaoChe {
		return
	}
	robPl, ok := attackObj.(player.Player)
	if !ok {
		return
	}

	biaoChe := n.(*biaochenpc.BiaocheNPC)
	transportationObj := biaoChe.GetTransportationObject()

	//劫个人镖计数
	if transportationObj.GetTransportType() != transportationtypes.TransportationTypeAlliance {
		robPlManager := robPl.GetPlayerDataManager(playertypes.PlayerTransportationType).(*transportationplayer.PlayerTransportationDataManager)
		robPlManager.RobTransportation(transportationObj.GetTransportId())
	}

	//被攻击信息推送
	pl := player.GetOnlinePlayerManager().GetPlayerById(transportationObj.GetPlayerId())
	if pl == nil {
		return
	}
	scTransportationProtectNotice := pbutil.BuildSCTransportationProtectNotice(defenceObj.GetPosition())
	pl.SendMsg(scTransportationProtectNotice)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectAttacked, event.EventListenerFunc(transportationBeAttack))
}
