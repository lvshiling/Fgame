package player

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"time"
)

const (
	taskTime = 3 * time.Second
)

type SongBuTingTask struct {
	p player.Player
}

func (wt *SongBuTingTask) Run() {
	manager := wt.p.GetPlayerDataManager(types.PlayerSongBuTingDataManagerType).(*PlayerSongBuTingDataManager)
	manager.refresh()
}

func (wt *SongBuTingTask) ElapseTime() time.Duration {
	return taskTime
}

func CreateSongBuTingTask(p player.Player) *SongBuTingTask {
	qTask := &SongBuTingTask{
		p: p,
	}
	return qTask
}
