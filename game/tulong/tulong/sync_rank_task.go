package tulong

import (
	"time"
)

const (
	syncTaskTime = time.Minute
)

type TuLongRankTask struct {
	s TuLongService
}

func (tt *TuLongRankTask) Run() {
	tt.s.SyncRemoteRankListTask()
}

func (tt *TuLongRankTask) ElapseTime() time.Duration {
	return syncTaskTime
}

func CreateSyncTask(s TuLongService) *TuLongRankTask {
	tuLongRankTask := &TuLongRankTask{
		s: s,
	}
	return tuLongRankTask
}
