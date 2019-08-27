package player

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"
)

const (
	checkTaskTime = 5 * time.Second
)

type DaLiWanTask struct {
	pl player.Player
}

func (ft *DaLiWanTask) Run() {
	manager := ft.pl.GetPlayerDataManager(playertypes.PlayerDaLiWanDataManagerType).(*PlayerDaLiWanManager)
	manager.CheckExpire()
}

func (ft *DaLiWanTask) ElapseTime() time.Duration {
	return checkTaskTime
}

func CreateDaLiWanTask(p player.Player) *DaLiWanTask {
	daLiWanTask := &DaLiWanTask{
		pl: p,
	}
	return daLiWanTask
}
