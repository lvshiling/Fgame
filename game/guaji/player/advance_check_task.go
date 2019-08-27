package player

import (
	gameevent "fgame/fgame/game/event"
	guajieventtypes "fgame/fgame/game/guaji/event/types"
	"fgame/fgame/game/player"
	"time"
)

const (
	advanceCheckTaskTime = time.Second * 10
)

type advanceCheckTask struct {
	p player.Player
}

func (t *advanceCheckTask) Run() {
	gameevent.Emit(guajieventtypes.GuaJiEventTypeGuaJiAdvanceCheck, t.p, nil)
	return
}

func (t *advanceCheckTask) ElapseTime() time.Duration {
	return advanceCheckTaskTime
}

func CreateAdvanceCheckTask(p player.Player) *advanceCheckTask {
	t := &advanceCheckTask{
		p: p,
	}
	return t
}
