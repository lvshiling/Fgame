package shenmo

import (
	"time"
)

type ShenMoRankTimeTask struct {
	s *shenMoService
}

const (
	timeTaskTime = time.Second * 3
)

func (t *ShenMoRankTimeTask) Run() {
	t.s.checkRefreshWeekRank()
}

//间隔时间
func (t *ShenMoRankTimeTask) ElapseTime() time.Duration {
	return timeTaskTime
}

func CreateShenMoRankTimeTask(s *shenMoService) *ShenMoRankTimeTask {
	t := &ShenMoRankTimeTask{
		s: s,
	}
	return t
}
