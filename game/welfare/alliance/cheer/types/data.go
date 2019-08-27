package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//城战助威
type AllianceCheerInfo struct {
	CheerGoldNum int32 `json:"goldNum"` //累计助威元宝
	IsEmail      bool  `json:"isEmail"` //是否邮件
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeAlliance, (*AllianceCheerInfo)(nil))
}
