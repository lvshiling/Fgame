package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/funcopen/funcopen"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/outlandboss/pbutil"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//外域boss重生
func battleObjectReborn(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	//校验怪物类型
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeOutlandBoss {
		return
	}

	funcOpenTemp := funcopen.GetFuncOpenService().GetFuncOpenTemplate(funcopentypes.FuncOpenTypeOutlandBoss)
	scMsg := pbutil.BuildSCOutlandBossInfoBroadcast(n)
	playerList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range playerList {
		if pl.GetLevel() < funcOpenTemp.OpenedLevel {
			continue
		}
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(battleObjectReborn))
}
