package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"
)

const (
	fashionTaskTime = 2 * time.Second
)

type FashionTask struct {
	pl player.Player
}

func (ft *FashionTask) Run() {
	now := global.GetGame().GetTimeService().Now()
	manager := ft.pl.GetPlayerDataManager(playertypes.PlayerFashionDataManagerType).(*PlayerFashionDataManager)
	expireMap := manager.GetExpireMap()
	if expireMap == nil {
		return
	}

	for _, fashionObj := range expireMap {
		_, activeFlag := manager.fashionRefreshCheck(fashionObj, now)
		if activeFlag {
			continue
		}
		manager.RemoveExpireFashion(fashionObj.FashionId)
	}
}

func (ft *FashionTask) ElapseTime() time.Duration {
	return fashionTaskTime
}

func CreateFashionTask(p player.Player) *FashionTask {
	fashionTask := &FashionTask{
		pl: p,
	}
	return fashionTask
}
