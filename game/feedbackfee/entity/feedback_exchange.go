package entity

//玩家逆付费数据
type FeedbackExchangeEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	PlayerId    int64  `gorm:"column:playerId"`
	ExchangeId  int64  `gorm:"column:exchangeId"`
	Code        string `gorm:"column:code"`
	Money       int32  `gorm:"column:money"`
	Status      int32  `gorm:"column:status"`
	ExpiredTime int64  `gorm:"column:expiredTime"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (e *FeedbackExchangeEntity) GetId() int64 {
	return e.Id
}

func (e *FeedbackExchangeEntity) TableName() string {
	return "t_feedback_exchange"
}
