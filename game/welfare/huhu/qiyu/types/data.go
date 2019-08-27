package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 奇遇副本
type HuHuQiYuInfo struct {
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeQiYu, (*HuHuQiYuInfo)(nil))
}
