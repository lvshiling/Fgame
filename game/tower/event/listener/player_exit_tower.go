package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/tower/pbutil"
	playertower "fgame/fgame/game/tower/player"
)

//玩家退出打宝塔
func playerExitScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	if !pl.GetScene().MapTemplate().IsTower() {
		return
	}

	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	towerManager.EndDaBao()

	totalExp := pl.GetCountTowerExp()
	dropMap := pl.GetCountTowerItemMap()
	remainTime := towerManager.GetRemainTime()
	scMsg := pbutil.BuildSCTowerResultNotice(totalExp, remainTime, dropMap)
	pl.SendMsg(scMsg)

	pl.ResetCountTower()
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(playerExitScene))
}
