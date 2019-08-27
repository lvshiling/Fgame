package trade

import (
	"time"
)

const (
	tradeRetryTaskTime = 5 * time.Second
)

type tradeRetryTask struct {
	s TradeService
}

func (t *tradeRetryTask) Run() {
	t.s.SyncRetryTradeList()
	return

}

func (t *tradeRetryTask) ElapseTime() time.Duration {
	return tradeRetryTaskTime
}

func createTradeRetryTask(s TradeService) *tradeRetryTask {
	t := &tradeRetryTask{
		s: s,
	}
	return t
}
