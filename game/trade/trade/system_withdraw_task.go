package trade

import (
	"time"
)

const (
	systemWithdrawTaskTime = time.Minute
)

type systemWithDrawTask struct {
	s TradeService
}

func (t *systemWithDrawTask) Run() {
	t.s.SystemWithdrawTradeList()
	return

}

func (t *systemWithDrawTask) ElapseTime() time.Duration {
	return systemWithdrawTaskTime
}

func createSystemWithDrawTask(s TradeService) *systemWithDrawTask {
	t := &systemWithDrawTask{
		s: s,
	}
	return t
}
