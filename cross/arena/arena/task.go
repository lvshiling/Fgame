package arena

import (
	"time"
)

const (
	matchTaskTime = time.Second * 3
)

type arenaMatchTask struct {
	s ArenaService
}

func (t *arenaMatchTask) Run() {
	t.s.Match()
}

func (t *arenaMatchTask) ElapseTime() time.Duration {
	return matchTaskTime
}

func CreateMatchTask(s ArenaService) *arenaMatchTask {
	t := &arenaMatchTask{
		s: s,
	}
	return t
}
