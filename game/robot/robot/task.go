package robot

import (
	"fgame/fgame/game/scene/scene"
	"time"
)

const (
	robotTaskTime = 300 * time.Millisecond
)

type RobotTask struct {
	p scene.RobotPlayer
}

func (t *RobotTask) Run() {
	t.p.GetCurrentAction().Action(t.p)
}

func (t *RobotTask) ElapseTime() time.Duration {
	return robotTaskTime
}

func CreateRobotTask(p scene.RobotPlayer) *RobotTask {
	t := &RobotTask{
		p: p,
	}
	return t
}
