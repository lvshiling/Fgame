package hongbao

import (
	"time"
)

type hongBaoExpireTask struct {
	s HongBaoService
}

func (t *hongBaoExpireTask) Run() {
	t.s.CheckExpireHongBao()
	return
}

var (
	hongBaoTaskTime = time.Minute * 5
)

func (t *hongBaoExpireTask) ElapseTime() time.Duration {
	return hongBaoTaskTime
}

func CreateHongBaoExpireTask(s HongBaoService) *hongBaoExpireTask {
	t := &hongBaoExpireTask{
		s: s,
	}
	return t
}
