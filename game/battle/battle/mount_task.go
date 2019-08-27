package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	"time"
)

//TODO 控制同步时间
const (
	mountTaskTime = time.Second
)

type mountTask struct {
	p scene.Player
}

func (t *mountTask) Run() {
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerMountCheck, t.p, nil)
}

func (t *mountTask) ElapseTime() time.Duration {
	return mountTaskTime
}

func CreateMountTask(p scene.Player) *mountTask {
	t := &mountTask{
		p: p,
	}
	return t
}
