package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 虎虎生风-采集物
type HuHuCollectInfo struct {
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeCollect, (*HuHuCollectInfo)(nil))
}
