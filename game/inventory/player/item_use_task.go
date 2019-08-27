package player

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"time"
)

const (
	resetItemUseTaskTime = time.Second * 5
)

// 刷新物品使用次数
type ItemUseTask struct {
	pl player.Player
}

func (t *ItemUseTask) Run() {
	inventoryManager := t.pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*PlayerInventoryDataManager)
	inventoryManager.refreshItemUseTimes()
}

func (t *ItemUseTask) ElapseTime() time.Duration {
	return resetItemUseTaskTime
}

func CreateResetItemUseTask(pl player.Player) *ItemUseTask {
	t := &ItemUseTask{
		pl: pl,
	}
	return t
}
