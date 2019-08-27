package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 炼制
type MadeInfo struct {
	Times int32 `json:"times"` //炼制次数
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeMade, welfaretypes.OpenActivityMadeSubTypeResource, (*MadeInfo)(nil))
}
