package battle

import (
	"fgame/fgame/game/scene/scene"
	"time"
)

//TODO 控制同步时间
const (
	reliveTaskTime = time.Second * 1
)

type reliveTask struct {
	p scene.Player
}

func (t *reliveTask) Run() {
	flag := t.p.RefreshReliveTime()
	if !flag {
		return
	}

}

func (t *reliveTask) ElapseTime() time.Duration {
	return reliveTaskTime
}

func CreateReliveTask(p scene.Player) *reliveTask {
	t := &reliveTask{
		p: p,
	}
	return t
}
