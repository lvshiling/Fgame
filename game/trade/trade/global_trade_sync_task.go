package trade

import (
	"time"
)

const (
	globalTradeSyncTaskTime = 5 * time.Second
)

type globalTradeSyncTask struct {
	s TradeService
}

func (t *globalTradeSyncTask) Run() {
	t.s.SyncGlobalTradeList()
	return

}

func (t *globalTradeSyncTask) ElapseTime() time.Duration {
	return globalTradeSyncTaskTime
}

func createGlobalTradeSyncTask(s TradeService) *globalTradeSyncTask {
	t := &globalTradeSyncTask{
		s: s,
	}
	return t
}
