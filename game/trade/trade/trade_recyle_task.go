package trade

import (
	"time"
)

const (
	tradeRecycleTaskTime = 10 * time.Minute
)

type tradeRecycleTask struct {
	s TradeService
}

func (t *tradeRecycleTask) Run() {
	t.s.SystemRecycle()
	return

}

func (t *tradeRecycleTask) ElapseTime() time.Duration {
	return tradeRecycleTaskTime
}

func createTradeRecycleTask(s TradeService) *tradeRecycleTask {
	t := &tradeRecycleTask{
		s: s,
	}
	return t
}
