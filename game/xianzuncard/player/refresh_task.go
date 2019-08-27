package player

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"time"
)

const (
	taskNoticeTime = time.Second * 5
)

type XianZunCardTask struct {
	pl player.Player
}

func (t *XianZunCardTask) Run() {
	manager := t.pl.GetPlayerDataManager(types.PlayerXianZunCardManagerType).(*PlayerXianZunCardDataManager)
	manager.refreshData()
}

func (t *XianZunCardTask) ElapseTime() time.Duration {
	return taskNoticeTime
}

func CreateXianZunCardTask(pl player.Player) *XianZunCardTask {
	xianZun := &XianZunCardTask{
		pl: pl,
	}
	return xianZun
}
