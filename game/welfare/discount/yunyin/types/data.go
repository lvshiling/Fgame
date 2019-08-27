package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type YunYinInfo struct {
	GoldNum       int32           `json:"goldNum"`
	ReceiveRecord []int32         `json:"receiveRecord"`
	BuyRecord     map[int32]int32 `json:"buyRecord"`
	IsEmail       bool            `json:"isEmail"`
}

// 是否能领取该档次奖励
func (y *YunYinInfo) IsCanReceive(goldNum int32) bool {
	if y.GoldNum < goldNum {
		return false
	}
	return true
}

// 是否已经领取了档次奖励
func (y *YunYinInfo) IsAlreadyReceive(goldNum int32) bool {
	for _, num := range y.ReceiveRecord {
		if num == goldNum {
			return true
		}
	}
	return false
}

// 添加购买记录
func (y *YunYinInfo) AddBuyRecord(typ int32, num int32) {
	_, ok := y.BuyRecord[typ]
	if !ok {
		y.BuyRecord[typ] = num
	} else {
		y.BuyRecord[typ] += num
	}
}

// 添加花费金额
func (y *YunYinInfo) AddGoldNum(goldNum int32) {
	y.GoldNum += goldNum
}

// 添加领取记录
func (y *YunYinInfo) AddReceiveRecord(goldNum int32) {
	y.ReceiveRecord = append(y.ReceiveRecord, goldNum)
}

// 获取购买数量
func (y *YunYinInfo) GetBuyNum(typ int32) int32 {
	num, ok := y.BuyRecord[typ]
	if !ok {
		return 0
	}
	return num
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeYunYin, (*YunYinInfo)(nil))
}
