package arena

import (
	"time"
)

const (
	rankTaskTime = time.Second * 5
)

type arenaRankTask struct {
	s ArenaService
}

func (t *arenaRankTask) Run() {
	t.s.CheckRefreshWeekRank()
}

func (t *arenaRankTask) ElapseTime() time.Duration {
	return rankTaskTime
}

func CreateArenaRankTask(s ArenaService) *arenaRankTask {
	t := &arenaRankTask{
		s: s,
	}
	return t
}
