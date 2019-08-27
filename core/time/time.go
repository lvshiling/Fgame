package global

import (
	"time"
)

type TimeService interface {
	Now() int64
	SetOffTime(offTime int64)
}

type timeService struct {
	offTime int64
}

func (ts *timeService) Now() int64 {
	return time.Now().UnixNano()/int64(time.Millisecond) + ts.offTime
}

func (ts *timeService) SetOffTime(offTime int64) {
	ts.offTime = offTime
}

func NewTimeService() TimeService {
	ts := &timeService{}
	return ts
}
