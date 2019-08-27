package welfare

import (
	"time"
)

type openTimeChangedTask struct {
	s *welfareService
}

func (t *openTimeChangedTask) Run() {
	t.s.checkopenTimeChanged()
	return
}

var (
	openTimeChangedTaskTime = time.Second * 5
)

func (t *openTimeChangedTask) ElapseTime() time.Duration {
	return openTimeChangedTaskTime
}

func CreateOpenTimeChangedTask(s *welfareService) *openTimeChangedTask {
	t := &openTimeChangedTask{
		s: s,
	}
	return t
}
