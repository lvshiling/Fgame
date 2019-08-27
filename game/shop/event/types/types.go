package types

import (
	commonlog "fgame/fgame/common/log"
)

type ShopEventType string

const (
	//商店事件
	EventTypeShopBuyItemLog ShopEventType = "ShopBuyItem"
)

type PlayerShopBuyItemLogEventData struct {
	shopId     int32
	buyNum     int32
	costMoney  int32
	reason     commonlog.ShopLogReason
	reasonText string
}

func CreatePlayerShopBuyItemLogEventData(shopId, buyNum, costMoney int32, reason commonlog.ShopLogReason, reasonText string) *PlayerShopBuyItemLogEventData {
	d := &PlayerShopBuyItemLogEventData{
		shopId:     shopId,
		buyNum:     buyNum,
		costMoney:  costMoney,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerShopBuyItemLogEventData) GetShopId() int32 {
	return d.shopId
}

func (d *PlayerShopBuyItemLogEventData) GetBuyNum() int32 {
	return d.buyNum
}

func (d *PlayerShopBuyItemLogEventData) GetCostMoney() int32 {
	return d.costMoney
}

func (d *PlayerShopBuyItemLogEventData) GetReason() commonlog.ShopLogReason {
	return d.reason
}

func (d *PlayerShopBuyItemLogEventData) GetReasonText() string {
	return d.reasonText
}
