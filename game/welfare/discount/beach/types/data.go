package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type BeachShopInfo struct {
	IsActivite int32           `json:"isActivite"` // 是否激活
	BuyRecord  map[int32]int32 `json:"buyRecord"`  // 购买商品记录
}

// 是否激活
func (b *BeachShopInfo) IsActivited() bool {
	return b.IsActivite != 0
}

// 激活成功
func (b *BeachShopInfo) ActiviteSuccess() {
	b.IsActivite = 1
}

func (b *BeachShopInfo) AddBuyRecord(typ int32, num int32) {
	_, ok := b.BuyRecord[typ]
	if !ok {
		b.BuyRecord[typ] = num
	} else {
		b.BuyRecord[typ] += num
	}
}

// 获取购买数量
func (b *BeachShopInfo) GetBuyNum(typ int32) int32 {
	num, ok := b.BuyRecord[typ]
	if !ok {
		return 0
	}

	return num
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeBeach, (*BeachShopInfo)(nil))
}
