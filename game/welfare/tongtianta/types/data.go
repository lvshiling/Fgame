package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 通天塔活动奖励数据
type TongTianTaInfo struct {
	MinForce  int32   `json:"minForce"`
	MaxForce  int32   `json:"maxForce"`
	ChargeNum int32   `json:"chargeNum"`
	Record    []int32 `json:"record"` // 领取信息
	IsEmail   bool    `json:"isEmail"`
}

// 充值的元宝是否足够
func (t *TongTianTaInfo) IsEnoughChargeNum(goldNum int32) bool {
	if t.ChargeNum < goldNum {
		return false
	}
	return true
}

// 是否已经领取过奖励
func (t *TongTianTaInfo) IsAlreadyReceiveByForce(force int32) bool {
	for _, info := range t.Record {
		if info == force {
			return true
		}
	}
	return false
}

// 领取奖励
func (t *TongTianTaInfo) ReceiveSuccess(force int32) bool {
	for _, info := range t.Record {
		if info == force {
			return false
		}
	}
	t.Record = append(t.Record, force)
	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeLingTong, (*TongTianTaInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeMingGe, (*TongTianTaInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeTuLong, (*TongTianTaInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeShengHen, (*TongTianTaInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeZhenFa, (*TongTianTaInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeBaby, (*TongTianTaInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeDianXing, (*TongTianTaInfo)(nil))
}
