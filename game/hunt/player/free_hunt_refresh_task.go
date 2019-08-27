package player

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"
)

const (
	taskNoticeTime = time.Second * 5
)

//免费寻宝跨天刷新
type FreeHuntRefreshTask struct {
	pl player.Player
}

func (t *FreeHuntRefreshTask) Run() {
	huntManager := t.pl.GetPlayerDataManager(playertypes.PlayerHuntDataManagerType).(*PlayerHuntDataManager)
	huntManager.RefreshFreeHuntTimes()
}

//间隔时间
func (t *FreeHuntRefreshTask) ElapseTime() time.Duration {
	return taskNoticeTime
}

func CreateFreeHuntRefresh(pl player.Player) *FreeHuntRefreshTask {
	t := &FreeHuntRefreshTask{
		pl: pl,
	}
	return t
}
