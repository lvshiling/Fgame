package types

import (
	centertypes "fgame/fgame/game/center/types"
)

//会员类型
type HuiYuanType int32

const (
	HuiYuanTypeCommon  HuiYuanType = iota // 普通玩家
	HuiYuanTypeInterim                    // 临时会员
	HuiYuanTypePlus                       // 至尊会员
)

func (t HuiYuanType) Valid() bool {
	switch t {
	case
		HuiYuanTypeCommon,
		HuiYuanTypeInterim,
		HuiYuanTypePlus:
		return true
	default:
		return false
	}
}

const (
	OnlineHoutaiType = centertypes.ZhiZunTypeExp
)
