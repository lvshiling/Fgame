package tower

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type towerDummyLogTask struct {
	sd *towerSceneData
}

func (t *towerDummyLogTask) Run() {
	err := t.sd.addDummyLog()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("tower:生成打宝塔虚拟日志错误")
		return
	}

	return
}

var (
	openTaskTime = time.Second * 10
)

func (t *towerDummyLogTask) ElapseTime() time.Duration {
	return openTaskTime
}

func CreateTowerDummyLogTask(sd *towerSceneData) *towerDummyLogTask {
	t := &towerDummyLogTask{
		sd: sd,
	}
	return t
}
