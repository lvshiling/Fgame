package activity

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type activityNoticeTask struct {
	as ActivityService
}

func (t *activityNoticeTask) Run() {
	err := t.as.CheckStartActivity()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("activity:活动开始提醒,错误")
		return
	}

	err = t.as.CheckEndActivity()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("activity:活动结束提醒,错误")
		return
	}
	return
}

var (
	noticeTaskTime = time.Second * 15
)

func (t *activityNoticeTask) ElapseTime() time.Duration {
	return noticeTaskTime
}

func CreateActivityNoticeTask(as ActivityService) *activityNoticeTask {
	t := &activityNoticeTask{
		as: as,
	}
	return t
}
