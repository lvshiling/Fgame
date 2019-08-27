package player

import (
	gameevent "fgame/fgame/game/event"
	playereventtypes "fgame/fgame/game/player/event/types"
	"time"
)

const (
	taskTime = time.Second * 5
)

//在线时间改变
type OnlineTimenChangedTask struct {
	pl *Player
}

func (t *OnlineTimenChangedTask) Run() {
	gameevent.Emit(playereventtypes.EventTypePlayerOnlineTimeChanged, t.pl, nil)
}

//间隔时间
func (t *OnlineTimenChangedTask) ElapseTime() time.Duration {
	return taskTime
}

func CreateOnlineTimeChangedTask(pl *Player) *OnlineTimenChangedTask {
	t := &OnlineTimenChangedTask{
		pl: pl,
	}
	return t
}
