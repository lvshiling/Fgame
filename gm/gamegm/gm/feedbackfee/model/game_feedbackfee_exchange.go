package model

type GamePlayerFeedBackFeeExchange struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	PlayerId    int64  `gorm:"column:playerId"`
	ExchangeId  int64  `gorm:"column:exchangeId"`
	ExpiredTime int64  `gorm:"column:expiredTime"`
	Money       int32  `gorm:"column:money"`
	Code        string `gorm:"column:code"`
	Status      int32  `gorm:"column:status"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (m *GamePlayerFeedBackFeeExchange) TableName() string {
	return "t_feedback_exchange"
}
