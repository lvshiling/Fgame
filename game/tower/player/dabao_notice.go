package player

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"
)

type towerNoticeTask struct {
	pl player.Player
}

func (t *towerNoticeTask) Run() {
	towerManager := t.pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*PlayerTowerDataManager)
	towerManager.noticeDaBaoTime()
	return
}

var (
	noticeTime = time.Minute
)

func (t *towerNoticeTask) ElapseTime() time.Duration {
	return noticeTime
}

func CreateNoticeTask(pl player.Player) *towerNoticeTask {
	t := &towerNoticeTask{
		pl: pl,
	}
	return t
}
