package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//砸金蛋
type SmashEggInfo struct {
	AttendTimes     int32   `json:"attendTimes"`     //已参与次数
	AttendTimesList []int32 `json:"attendTimesList"` //已砸金蛋记录
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmashEgg, (*SmashEggInfo)(nil))
}
