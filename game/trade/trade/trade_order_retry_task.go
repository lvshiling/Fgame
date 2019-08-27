package trade

import (
	"time"
)

const (
	orderRetryTaskTime = time.Minute
)

type tradeOrderRetryTask struct {
	s TradeService
}

func (t *tradeOrderRetryTask) Run() {
	t.s.SyncRetryOrderList()
	return

}

func (t *tradeOrderRetryTask) ElapseTime() time.Duration {
	return orderRetryTaskTime
}

func createTradeOrderRetryTask(s TradeService) *tradeOrderRetryTask {
	t := &tradeOrderRetryTask{
		s: s,
	}
	return t
}
