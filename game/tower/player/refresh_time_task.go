package player

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"
)

type towerRefreshTimeTask struct {
	pl player.Player
}

func (t *towerRefreshTimeTask) Run() {
	towerManager := t.pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*PlayerTowerDataManager)
	towerManager.checkTowerTime()
	return
}

var (
	openTaskTime = time.Second * 10
)

func (t *towerRefreshTimeTask) ElapseTime() time.Duration {
	return openTaskTime
}

func CreateRefreshTimeTask(pl player.Player) *towerRefreshTimeTask {
	t := &towerRefreshTimeTask{
		pl: pl,
	}
	return t
}
