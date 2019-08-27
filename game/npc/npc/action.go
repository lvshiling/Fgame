package npc

import (
	"fgame/fgame/game/scene/scene"
	"time"
)

const (
	actionElapseTime = 300 * time.Millisecond
)

type npcTask struct {
	n scene.NPC
}

func (t *npcTask) Run() {
	currentAction := t.n.GetCurrentAction()
	if currentAction == nil {
		return
	}
	currentAction.Tick(t.n)
}

func (t *npcTask) ElapseTime() time.Duration {
	return actionElapseTime
}

func CreateNPCTask(n scene.NPC) *npcTask {
	t := &npcTask{
		n: n,
	}
	return t
}
