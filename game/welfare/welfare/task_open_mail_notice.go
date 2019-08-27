package welfare

import (
	"time"
)

type startEmailTask struct {
	s *welfareService
}

func (t *startEmailTask) Run() {
	t.s.checkActivityStartMail()
	return
}

var (
	startEmailTaskTime = time.Second * 5
)

func (t *startEmailTask) ElapseTime() time.Duration {
	return startEmailTaskTime
}

func CreateStartEmailTask(s *welfareService) *startEmailTask {
	t := &startEmailTask{
		s: s,
	}
	return t
}
