package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	"time"
)

//TODO 控制同步时间
const (
	pkTaskTime = time.Second
)

type pkTask struct {
	p scene.Player
}

func (t *pkTask) Run() {
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerPkCheck, t.p, nil)
}

func (t *pkTask) ElapseTime() time.Duration {
	return pkTaskTime
}

func CreatePkTask(p scene.Player) *pkTask {
	t := &pkTask{
		p: p,
	}
	return t
}
