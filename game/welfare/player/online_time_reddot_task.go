package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	"time"
)

const (
	taskNoticeTime = time.Second * 30
)

//运营活动-时间改变红点推送
type TimeReddotNoticeTask struct {
	pl player.Player
}

func (t *TimeReddotNoticeTask) Run() {
	gameevent.Emit(welfareeventtypes.EventTypeTimeReddotNotice, t.pl, nil)
}

//间隔时间
func (t *TimeReddotNoticeTask) ElapseTime() time.Duration {
	return taskNoticeTime
}

func CreateTimeReddotNoticeTask(pl player.Player) *TimeReddotNoticeTask {
	t := &TimeReddotNoticeTask{
		pl: pl,
	}
	return t
}
