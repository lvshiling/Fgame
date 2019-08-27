package chuangshi

import (
	"time"
)

const (
	cityTaskTime = time.Minute * 1
)

type cityTask struct {
	s ChuangShiService
}

func (t *cityTask) Run() {
	t.s.CityTask()
}

func (t *cityTask) ElapseTime() time.Duration {
	return cityTaskTime
}

func CreateCityTask(s ChuangShiService) *cityTask {
	t := &cityTask{
		s: s,
	}
	return t
}
