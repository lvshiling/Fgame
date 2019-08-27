package model

type CenterPlayerFeedBackFeeExchange struct {
	Id          int64  `gorm:"primary_key;column:id"`
	Platform    int32  `gorm:"column:platform"`
	ServerId    int32  `gorm:"column:serverId"`
	PlayerId    int64  `gorm:"column:playerId"`
	ExchangeId  int64  `gorm:"column:exchangeId"`
	ExpiredTime int64  `gorm:"column:expiredTime"`
	Money       int32  `gorm:"column:money"`
	Code        string `gorm:"column:code"`
	Status      int32  `gorm:"column:status"`
	WxId        string `gorm:"column:wxId"`
	OrderId     string `gorm:"column:orderId"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (m *CenterPlayerFeedBackFeeExchange) TableName() string {
	return "t_feedbackfee_exchange"
}
