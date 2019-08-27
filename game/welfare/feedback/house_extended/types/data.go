package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-房产活动
type FeedbackHouseExtendedInfo struct {
	ActivateChargeNum   int32 `json:"activateChargeNum"`   //激活礼包充值数
	IsActivateGift      bool  `json:"isActivateGift"`      //是否领取激活礼包
	UplevelChargeNum    int32 `json:"uplevelChargeNum"`    //装修礼包充值数
	CurUplevelGiftLevel int32 `json:"curUplevelGiftLevel"` //当前装修礼包等级
	IsUplevelGift       bool  `json:"isUplevelGift"`       //是否领取装修礼包
}

//充值元宝
func (info *FeedbackHouseExtendedInfo) AddChargeNum(goldNum int32) {
	info.ActivateChargeNum += goldNum
	info.UplevelChargeNum += goldNum
}

//是否能领取激活礼包
func (info *FeedbackHouseExtendedInfo) IsCanReceiveActivateGift(needCharge int32) bool {
	if info.IsActivateGift {
		return false
	}

	return info.ActivateChargeNum >= needCharge
}

//领取激活礼包
func (info *FeedbackHouseExtendedInfo) ReceiveActivateGift() {
	info.IsActivateGift = true
}

//是否能领取装修礼包
func (info *FeedbackHouseExtendedInfo) IsCanReceiveUplevelGift(needCharge int32) bool {
	if info.IsUplevelGift {
		return false
	}

	return info.UplevelChargeNum >= needCharge
}

//领取装修礼包
func (info *FeedbackHouseExtendedInfo) ReceiveUplevelGift() {
	info.IsUplevelGift = true
}

//装修礼包跨天
func (info *FeedbackHouseExtendedInfo) CrossDayUplevelGift() {
	if info.IsUplevelGift {
		info.IsUplevelGift = false
		info.CurUplevelGiftLevel += 1
	}

	info.UplevelChargeNum = 0
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseExtended, (*FeedbackHouseExtendedInfo)(nil))
}
