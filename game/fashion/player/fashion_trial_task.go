package player

import (
	fashiontypes "fgame/fgame/game/fashion/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"
)

const (
	fashionTrialTaskTime = 5 * time.Second
)

type FashionTrialTask struct {
	pl player.Player
}

func (ft *FashionTrialTask) Run() {
	now := global.GetGame().GetTimeService().Now()
	manager := ft.pl.GetPlayerDataManager(playertypes.PlayerFashionDataManagerType).(*PlayerFashionDataManager)

	//时装试用卡
	trialMap := manager.GetTrialFashionMap()
	for _, trial := range trialMap {
		if trial.expireTime == 0 {
			continue
		}

		if now < trial.expireTime {
			continue
		}

		manager.TrialFashionOverdue(trial.trialFashionId, fashiontypes.FashionTrialOverdueTypeExpire)
	}

}

func (ft *FashionTrialTask) ElapseTime() time.Duration {
	return fashionTrialTaskTime
}

func CreateFashionTrialTask(p player.Player) *FashionTrialTask {
	fashionTask := &FashionTrialTask{
		pl: p,
	}
	return fashionTask
}
