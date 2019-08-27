package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//限时折扣-转生大礼包
type DiscountZhuanShengInfo struct {
	BuyRecord         map[int32]int32 `json:"buyRecord"`         //购买记录
	GiftReceiveRecord map[int32]int32 `json:"giftReceiveRecord"` //赠品领取记录
	ChargeNum         int64           `json:"chargeNum"`         //充值数
	UsePoint          int32           `json:"usePoint"`          //已使用积分
}

func (info *DiscountZhuanShengInfo) IsBuy(giftType int32) bool {
	_, ok := info.BuyRecord[giftType]
	if !ok {
		return false
	}
	return true
}

func (info *DiscountZhuanShengInfo) AddGiftRecord(giftType int32) {
	_, ok := info.GiftReceiveRecord[giftType]
	if !ok {
		info.GiftReceiveRecord[giftType] = 1
	} else {
		info.GiftReceiveRecord[giftType] += 1
	}
}

func (info *DiscountZhuanShengInfo) IsCanReceiveGift(giftType int32) bool {
	if !info.IsBuy(giftType) {
		return false
	}

	_, ok := info.GiftReceiveRecord[giftType]
	if ok {
		return false
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeZhuanSheng, (*DiscountZhuanShengInfo)(nil))
}
