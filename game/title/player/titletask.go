package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"time"
)

const (
	activityTaskTime = 2 * time.Second
)

type ActivityTitleTask struct {
	p player.Player
}

func (att *ActivityTitleTask) Run() {
	now := global.GetGame().GetTimeService().Now()
	manager := att.p.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*PlayerTitleDataManager)
	activityMap := manager.GetActivityMap()
	if activityMap == nil {
		return
	}

	for _, activityObj := range activityMap {
		_, activeFlag := manager.titleRefreshCheck(activityObj, now)
		if activeFlag {
			continue
		}
		manager.RemoveActivity(activityObj.TitleId)
	}

}

func (att *ActivityTitleTask) ElapseTime() time.Duration {
	return activityTaskTime
}

func CreateActivityTitleTask(p player.Player) *ActivityTitleTask {
	titleTask := &ActivityTitleTask{
		p: p,
	}
	return titleTask
}
