package buff

import (
	"fgame/fgame/game/scene/scene"
	"time"
)

const (
	buffTaskTime = time.Second
)

type BuffTask struct {
	bo scene.BattleObject
}

func (t *BuffTask) Run() {
	t.bo.RefreshBuff()
}

func (t *BuffTask) ElapseTime() time.Duration {
	return buffTaskTime
}

func CreateBuffTask(bo scene.BattleObject) *BuffTask {
	t := &BuffTask{
		bo: bo,
	}
	return t
}
