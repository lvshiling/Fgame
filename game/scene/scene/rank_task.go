package scene

import (
	"time"
)

const (
	rankTaskTime = 3 * time.Second
)

type rankTask struct {
	s Scene
}

func (t *rankTask) Run() {
	t.s.Sort()
}

func (t *rankTask) ElapseTime() time.Duration {
	return rankTaskTime
}

func createRankTask(s Scene) *rankTask {
	t := &rankTask{
		s: s,
	}
	return t
}
