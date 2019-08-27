package alliance

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type douShenTask struct {
	s AllianceService
}

func (t *douShenTask) Run() {
	err := t.s.UpdateDouShenList()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("alliance:整点更新斗神列表,错误")
		return
	}
	return
}

var (
	douShenTaskTime = time.Minute
)

func (t *douShenTask) ElapseTime() time.Duration {
	return douShenTaskTime
}

func CreateDouShenTask(s AllianceService) *douShenTask {
	t := &douShenTask{
		s: s,
	}
	return t
}
