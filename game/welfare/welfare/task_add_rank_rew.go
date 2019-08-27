package welfare

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type ranRewTask struct {
	s *welfareService
}

func (t *ranRewTask) Run() {
	err := t.s.addOpenRankRewards()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("welfare:发送排行榜奖励邮件错误")
		return
	}

	return
}

var (
	addRewTaskTime = time.Minute
)

func (t *ranRewTask) ElapseTime() time.Duration {
	return addRewTaskTime
}

func CreateAddRankRewTask(s *welfareService) *ranRewTask {
	t := &ranRewTask{
		s: s,
	}
	return t
}
