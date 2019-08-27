package alliance

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type allianceSceneTask struct {
	s AllianceService
}

func (t *allianceSceneTask) Run() {
	//城战场景
	err := t.s.CheckAllianceScene()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("alliance:定时开启九霄城战,错误")
		return
	}
}

var (
	allianceSceneTaskTime = time.Minute
)

func (t *allianceSceneTask) ElapseTime() time.Duration {
	return allianceSceneTaskTime
}

func CreateAllianceSceneTask(s AllianceService) *allianceSceneTask {
	t := &allianceSceneTask{
		s: s,
	}
	return t
}
