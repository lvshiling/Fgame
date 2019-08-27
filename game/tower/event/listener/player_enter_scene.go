package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/tower/pbutil"
	playertower "fgame/fgame/game/tower/player"
	"fgame/fgame/game/tower/tower"
)

//玩家进入打宝塔
func playerEnterTower(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	s := pl.GetScene()
	if !s.MapTemplate().IsTower() {
		return
	}

	sd, ok := s.SceneDelegate().(tower.TowerSceneData)
	if !ok {
		return
	}

	towerTemp := sd.GetTowerTemplate()
	logList := sd.GetLogByTime(0)
	floor := int32(towerTemp.Id)
	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	remainTime := towerManager.GetRemainTime()
	n := tower.GetTowerService().GetTowerBoss(floor)
	scMsg := pbutil.BuildSCTowerEnter(n, remainTime, logList, floor)
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterTower))
}
