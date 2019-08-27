package types

import (
	"fgame/fgame/game/common/common"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//限时折扣-砍价礼包
type DiscountKanJiaInfo struct {
	GoldNum      int64                 `json:"goldNum"`      //充值元宝
	UseTimes     int32                 `json:"useTimes"`     //砍价次数
	BuyRecord    []int32               `json:"buyRecord"`    //购买记录
	KanJiaRecord map[int32]*KanJiaInfo `json:"kanJiaRecord"` //砍价记录
}

type KanJiaInfo struct {
	KanJiaTimes int32 `json:"kanJiaTimes"` //砍价次数
	Discount    int32 `json:"discount"`    //折扣
}

func (info *DiscountKanJiaInfo) AddBuyRecord(giftType int32) {
	info.BuyRecord = append(info.BuyRecord, int32(giftType))
}

func (info *DiscountKanJiaInfo) UpdateKanJiaRecord(giftType int32, discount int32) {
	kanJiaInfo, ok := info.KanJiaRecord[giftType]
	if !ok {
		kanJiaInfo = &KanJiaInfo{}
		info.KanJiaRecord[giftType] = kanJiaInfo
	}
	kanJiaInfo.Discount = discount
	kanJiaInfo.KanJiaTimes += 1
	info.UseTimes += 1
}

func (info *DiscountKanJiaInfo) GetKanJiaInfo(giftType int32) (int32, int32) {
	kanJiaInfo, ok := info.KanJiaRecord[giftType]
	if !ok {
		return 0, int32(common.MAX_RATE)
	}

	return kanJiaInfo.KanJiaTimes, kanJiaInfo.Discount
}

func (info *DiscountKanJiaInfo) CountRewTimes(ratio int32) int32 {
	if ratio == 0 {
		panic(fmt.Errorf("兑换系数和最小起赠次数不能为0, ratio:%d", ratio))
	}
	leftGoldNum := int32(info.GoldNum) - info.UseTimes*ratio
	return leftGoldNum / ratio
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeKanJia, (*DiscountKanJiaInfo)(nil))
}
