package lingtong

import (
	"fgame/fgame/game/scene/scene"
	"time"
)

const (
	actionElapseTime = 300 * time.Millisecond
)

type lingTongTask struct {
	l scene.LingTong
}

func (t *lingTongTask) Run() {

	owner := t.l.GetOwner()
	if !owner.IsRobot() && !owner.IsGuaJi() {
		return
	}
	currentAction := t.l.GetCurrentAction()
	if currentAction == nil {
		return
	}
	currentAction.Action(t.l)
}

func (t *lingTongTask) ElapseTime() time.Duration {
	return actionElapseTime
}

func CreateLingTongTask(l scene.LingTong) *lingTongTask {
	t := &lingTongTask{
		l: l,
	}
	return t
}
