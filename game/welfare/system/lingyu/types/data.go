package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//领域激活
type SystemLingYuInfo struct {
	StartTime           int64 `json:startTime`             //开启时间
	IsOpen              bool  `json:isOpen`                //是否过了开启时间
	IsActivate          bool  `json:"isActivate"`          //是否激活
	MaxSingleChargeGold int32 `json:"maxSingleChargeGold"` //最大单笔充值
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeSystemActivate, welfaretypes.OpenActivitySystemActivateSubTypeLingYu, (*SystemLingYuInfo)(nil))
}
