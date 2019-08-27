package exchange

import (
	couponservertypes "fgame/fgame/coupon_server/types"
)

type FeedbackfeeExchangeObject struct {
	id          int64
	platform    int32
	serverId    int32
	playerId    int64
	playerName  string
	playerLevel int32
	exchangeId  int64
	expiredTime int64
	money       int32
	code        string
	status      couponservertypes.ExchangeStatus
	wxId        string
	orderId     string
	createTime  int64
	deleteTime  int64
	updateTime  int64
}

func NewFeedbackfeeExchangeObject() *FeedbackfeeExchangeObject {
	o := &FeedbackfeeExchangeObject{}
	return o
}

func convertFeedbackfeeExchangeObjectToEntity(o *FeedbackfeeExchangeObject) (*FeedbackfeeExchangeEntity, error) {

	e := &FeedbackfeeExchangeEntity{
		Id:       o.id,
		Platform: o.platform,
		ServerId: o.serverId,
		PlayerId: o.playerId,

		ExchangeId:  o.exchangeId,
		ExpiredTime: o.expiredTime,
		Money:       o.money,
		Code:        o.code,
		Status:      int32(o.status),
		WxId:        o.wxId,
		OrderId:     o.orderId,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *FeedbackfeeExchangeObject) GetId() int64 {
	return o.id
}

func (o *FeedbackfeeExchangeObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *FeedbackfeeExchangeObject) GetPlatform() int32 {
	return o.platform
}

func (o *FeedbackfeeExchangeObject) GetServerId() int32 {
	return o.serverId
}
func (o *FeedbackfeeExchangeObject) GetExchangeId() int64 {
	return o.exchangeId
}

func (o *FeedbackfeeExchangeObject) GetMoney() int32 {
	return o.money
}

func (o *FeedbackfeeExchangeObject) GetCode() string {
	return o.code
}

func (o *FeedbackfeeExchangeObject) GetOrderId() string {
	return o.orderId
}

func (o *FeedbackfeeExchangeObject) GetWxId() string {
	return o.wxId
}
func (o *FeedbackfeeExchangeObject) GetStatus() couponservertypes.ExchangeStatus {
	return o.status
}

func (o *FeedbackfeeExchangeObject) ToEntity() (e *FeedbackfeeExchangeEntity, err error) {
	e, err = convertFeedbackfeeExchangeObjectToEntity(o)
	return e, err
}

func (o *FeedbackfeeExchangeObject) FromEntity(e *FeedbackfeeExchangeEntity) (err error) {
	o.id = e.Id
	o.platform = e.Platform
	o.serverId = e.ServerId
	o.playerId = e.PlayerId

	o.exchangeId = e.ExchangeId
	o.expiredTime = e.ExpiredTime
	o.money = e.Money
	o.code = e.Code
	o.status = couponservertypes.ExchangeStatus(e.Status)
	o.wxId = e.WxId
	o.orderId = e.OrderId
	o.updateTime = e.UpdateTime
	o.createTime = e.CreateTime
	o.deleteTime = e.DeleteTime
	return nil
}

func (o *FeedbackfeeExchangeObject) Expired(now int64) bool {
	if o.status != couponservertypes.ExchangeStatusInit {
		return false
	}
	o.status = couponservertypes.ExchangeStatusExpired
	o.updateTime = now
	return true
}

func (o *FeedbackfeeExchangeObject) Finish(wxId string, orderId string, now int64) bool {
	if o.status != couponservertypes.ExchangeStatusInit {
		return false
	}
	o.status = couponservertypes.ExchangeStatusFinish
	o.wxId = wxId
	o.orderId = orderId
	o.updateTime = now
	return true
}

func (o *FeedbackfeeExchangeObject) Notify(now int64) bool {
	if o.status != couponservertypes.ExchangeStatusFinish {
		return false
	}
	o.status = couponservertypes.ExchangeStatusNotify
	o.updateTime = now
	return true
}
