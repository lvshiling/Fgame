package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questeventtypes "fgame/fgame/game/quest/event/types"
	"time"
)

const (
	qiyuTaskTime = 30 * time.Second
)

type QiYuTask struct {
	p player.Player
}

func (t *QiYuTask) Run() {
	manager := t.p.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*PlayerQuestDataManager)
	now := global.GetGame().GetTimeService().Now()
	for _, qiyu := range manager.GetQiYuMap() {
		if now < qiyu.endTime {
			continue
		}

		gameevent.Emit(questeventtypes.EventTypeQuestQiYuEnd, t.p, qiyu)
	}
}

func (t *QiYuTask) ElapseTime() time.Duration {
	return qiyuTaskTime
}

func CreateQiYuTask(p player.Player) *QiYuTask {
	qTask := &QiYuTask{
		p: p,
	}
	return qTask
}
