package player

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"time"
)

const (
	wingTaskTime = 2 * time.Second
)

type WingTask struct {
	p player.Player
}

func (wt *WingTask) Run() {
	manager := wt.p.GetPlayerDataManager(types.PlayerWingDataManagerType).(*PlayerWingDataManager)
	overFlag := manager.WingTrialIsOverdued()
	if overFlag {
		manager.RemoveWingTrial(true)
	}
}

func (wt *WingTask) ElapseTime() time.Duration {
	return wingTaskTime
}

func CreateWingTask(p player.Player) *WingTask {
	wingTask := &WingTask{
		p: p,
	}
	return wingTask
}
