package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 虎虎生风-掉落类型
type HuHuInfo struct {
	CurDayDropNum int32 `json:"curDayDropNum"` //当日怪物掉落量
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeDrop, (*HuHuInfo)(nil))
}
