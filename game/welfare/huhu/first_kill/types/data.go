package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 虎虎生风-首杀
type HuHuFirstKillInfo struct {
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeFirstDrop, (*HuHuFirstKillInfo)(nil))
}
