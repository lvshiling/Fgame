package welfare

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type labaDummyLogTask struct {
	s *welfareService
}

func (t *labaDummyLogTask) Run() {
	// err := t.s.addDummyLaBaLog()
	// if err != nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"error": err,
	// 		}).Error("welfare:生成拉霸虚拟日志错误")
	// 	return
	// }

	err := t.s.addDummyDrewLog()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("welfare:生成抽奖虚拟日志错误")
		return
	}

	err = t.s.addDummyCrazyBoxLog()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("welfare:生成疯狂宝箱虚拟日志错误")
		return
	}

	return
}

var (
	addLogTaskTime = time.Second * 15
)

func (t *labaDummyLogTask) ElapseTime() time.Duration {
	return addLogTaskTime
}

func CreateLaBaDummyLogTask(s *welfareService) *labaDummyLogTask {
	t := &labaDummyLogTask{
		s: s,
	}
	return t
}
