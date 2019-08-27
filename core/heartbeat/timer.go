package heartbeat

import (
	coretime "fgame/fgame/core/time"
	"sync"
	"time"
)

var (
	once sync.Once
	ts   coretime.TimeService
)

//只初始化一次
func SetupTimeService(ats coretime.TimeService) {
	once.Do(func() {
		ts = ats
	})
}

type timer struct {
	d      time.Duration
	timeUp int64
}

func (t *timer) IsTimeUp() bool {
	//获取时间
	now := ts.Now()
	if t.timeUp <= now {
		return true
	}
	return false
}

func (t *timer) Reset(d time.Duration) {
	now := ts.Now()
	t.timeUp = now + int64(d/time.Millisecond)
}

//TODO 对象池
func createTimer(d time.Duration) *timer {
	t := &timer{}
	now := ts.Now()
	t.timeUp = now + int64(d/time.Millisecond)
	return t
}
