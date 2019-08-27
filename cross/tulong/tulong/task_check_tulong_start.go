package tulong

import (
	"time"
)

const (
	checkStartTaskTime = time.Second * 3
)

type TuLongStartTask struct {
	s TuLongService
}

func (tt *TuLongStartTask) Run() {
	tt.s.CheckTuLongActivityTask()
}

func (tt *TuLongStartTask) ElapseTime() time.Duration {
	return checkStartTaskTime
}

func CreateTuLongStartTask(s TuLongService) *TuLongStartTask {
	tuLongStartTask := &TuLongStartTask{
		s: s,
	}
	return tuLongStartTask
}
