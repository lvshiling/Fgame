package welfare

import (
	"time"
)

type refreshXunHuanTask struct {
	s *welfareService
}

func (t *refreshXunHuanTask) Run() {
	t.s.checkRefreshXunHuan()
	t.s.checkTempRank()
	return
}

var (
	refreshXunHuanTaskTime = time.Second * 5
)

func (t *refreshXunHuanTask) ElapseTime() time.Duration {
	return refreshXunHuanTaskTime
}

func CreateRefreshXunHuanTask(s *welfareService) *refreshXunHuanTask {
	t := &refreshXunHuanTask{
		s: s,
	}
	return t
}
