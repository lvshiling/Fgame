package chuangshi

import (
	"time"
)

const (
	shenWangTaskTime = time.Second * 10
)

type shenWangTask struct {
	s ChuangShiService
}

func (t *shenWangTask) Run() {
	t.s.ShenWangTask()
}

func (t *shenWangTask) ElapseTime() time.Duration {
	return shenWangTaskTime
}

func CreateShenWangTask(s ChuangShiService) *shenWangTask {
	t := &shenWangTask{
		s: s,
	}
	return t
}
