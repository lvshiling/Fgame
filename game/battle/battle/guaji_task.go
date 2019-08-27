package battle

import (
	"fgame/fgame/game/scene/scene"
	"time"
)

const (
	guaJiTaskTime = 300 * time.Millisecond
)

type GuaJiTask struct {
	p scene.Player
}

func (t *GuaJiTask) Run() {
	t.p.GetCurrentGuaJiAction().GuaJi(t.p)
}

func (t *GuaJiTask) ElapseTime() time.Duration {
	return guaJiTaskTime
}

func CreateGuaJiTask(p scene.Player) *GuaJiTask {
	t := &GuaJiTask{
		p: p,
	}
	return t
}
