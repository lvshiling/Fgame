package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//炸矿活动
type BombOreInfo struct {
	AttendTimes int32 `json:"attendTimes"` //已参与次数
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeBombOre, (*BombOreInfo)(nil))
}
