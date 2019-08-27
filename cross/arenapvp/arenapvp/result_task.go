package arenapvp

import (
	"time"
)

const (
	resultTaskTime = time.Minute
)

type arenapvpResultTask struct {
	s ArenapvpService
}

func (t *arenapvpResultTask) Run() {
	t.s.RefreshBattleResult()
}

func (t *arenapvpResultTask) ElapseTime() time.Duration {
	return resultTaskTime
}

func CreateArenapvpResultTask(s ArenapvpService) *arenapvpResultTask {
	t := &arenapvpResultTask{
		s: s,
	}
	return t
}
