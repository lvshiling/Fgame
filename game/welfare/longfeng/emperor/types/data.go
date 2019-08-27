package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 龙凤呈祥
type LongFengInfo struct {
	RobTimes int32 `json:"robTimes"` //抢夺次数
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeLongFeng, welfaretypes.OpenActivityDefaultSubTypeDefault, (*LongFengInfo)(nil))
}
