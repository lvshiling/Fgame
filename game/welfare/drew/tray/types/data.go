package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//抽奖活动
type LuckyDrewInfo struct {
	AttendTimes int32 `json:"attendTimes"` //已参与次数
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeTray, (*LuckyDrewInfo)(nil))
}
