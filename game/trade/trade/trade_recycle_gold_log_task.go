package trade

import (
	"time"
)

const (
	logRecycleGoldTaskTime = time.Minute
)

type tradeRecycleGoldTask struct {
	s *tradeService
}

func (t *tradeRecycleGoldTask) Run() {
	t.s.sendRecycleGoldLog()
	return

}

func (t *tradeRecycleGoldTask) ElapseTime() time.Duration {
	return logRecycleGoldTaskTime
}

func createTradeRecycleGoldTask(s *tradeService) *tradeRecycleGoldTask {
	t := &tradeRecycleGoldTask{
		s: s,
	}
	return t
}
