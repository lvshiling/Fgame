package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/foe/foe"
	"fgame/fgame/game/foe/pbutil"
	playerfoe "fgame/fgame/game/foe/player"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//玩家死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.BattleObject)
	if !ok {
		return
	}
	pl, ok := bo.(player.Player)
	if !ok {
		return
	}

	attackId := data.(int64)
	foePl := player.GetOnlinePlayerManager().GetPlayerById(attackId)
	if foePl == nil {
		return
	}

	s := pl.GetScene()
	if s != foePl.GetScene() {
		return
	}

	// 仇人推送
	sceneType := s.MapTemplate().GetMapType()
	foe.FoeInfoNotice(pl, foePl, sceneType)

	// 添加仇人
	manager := pl.GetPlayerDataManager(types.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	flag := manager.AddFoe(attackId)
	if !flag {
		return
	}
	foeInfo, err := player.GetPlayerService().GetPlayerInfo(attackId)
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	tempFoe := pbutil.BuildFoe(attackId, now, foeInfo)
	scFoeAdd := pbutil.BuildSCFoeAdd(tempFoe)
	pl.SendMsg(scFoeAdd)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
