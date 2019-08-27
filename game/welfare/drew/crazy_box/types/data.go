package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//疯狂宝箱
type CrazyBoxInfo struct {
	GoldNum     int32 `json:"goldNum"`     //消费元宝
	AttendTimes int32 `json:"attendTimes"` //已参与次数
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCrazyBox, (*CrazyBoxInfo)(nil))
}
