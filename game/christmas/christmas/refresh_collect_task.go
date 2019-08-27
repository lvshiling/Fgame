package christmas

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type refreshCollectTask struct {
	s *christmasService
}

func (t *refreshCollectTask) Run() {
	err := t.s.checkRefreshCollect()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("christmas:圣诞采集定时刷新错误")
		return
	}

	return
}

var (
	refresTaskTime = time.Second * 5
)

func (t *refreshCollectTask) ElapseTime() time.Duration {
	return refresTaskTime
}

func CreateRefreshCollectTask(s *christmasService) *refreshCollectTask {
	t := &refreshCollectTask{
		s: s,
	}
	return t
}
