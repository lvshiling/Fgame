package scene

import (
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"time"
)

const (
	robot         = true
	robotTaskTime = time.Second * 5
)

type SceneRobotTask struct {
	s Scene
}

func (t *SceneRobotTask) Run() {
	if !robot {
		return
	}
	if !t.s.MapTemplate().IsWorld() {
		return
	}
	gameevent.Emit(sceneeventtypes.EventTypeSceneRobotCheck, t.s, nil)
}

func (t *SceneRobotTask) ElapseTime() time.Duration {
	return robotTaskTime
}

func CreateSceneRobotTask(s Scene) *SceneRobotTask {
	t := &SceneRobotTask{
		s: s,
	}
	return t
}
