package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/lingtong/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"
)

const (
	fashionTaskTime = 2 * time.Second
)

type LingTongFashionTask struct {
	pl player.Player
}

func (ft *LingTongFashionTask) Run() {
	now := global.GetGame().GetTimeService().Now()
	manager := ft.pl.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*PlayerLingTongDataManager)
	activateFashionMap := manager.GetActivateFashionMap()
	if len(activateFashionMap) != 0 {
		for fashionId, _ := range activateFashionMap {
			lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
			if lingTongFashionTemplate == nil {
				continue
			}
			if lingTongFashionTemplate.Time == 0 {
				continue
			}
			expireFlag := manager.fashionRefreshCheck(fashionId, now)
			if expireFlag {
				gameevent.Emit(lingtongeventtypes.EventTypeLingTongFashionOverdue, ft.pl, fashionId)
			}
		}
	}

	fashionTrialObj := manager.GetFashionTrialObject()
	if fashionTrialObj != nil && !fashionTrialObj.GetIsExpire() {
		expireTime := fashionTrialObj.GetActivateTime() + fashionTrialObj.GetDurationTime()
		if expireTime >= now {
			manager.TrialFashionOverdue(fashionTrialObj.GetTrialFashionId(), types.LingTongFashionTrialOverdueTypeExpire)
		}
	}

}

func (ft *LingTongFashionTask) ElapseTime() time.Duration {
	return fashionTaskTime
}

func CreateLingTongFashionTask(p player.Player) *LingTongFashionTask {
	fashionTask := &LingTongFashionTask{
		pl: p,
	}
	return fashionTask
}
