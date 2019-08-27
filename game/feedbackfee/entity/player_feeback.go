package entity

//玩家逆付费数据
type PlayerFeedbackFeeEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	TotalGetMoney int64 `gorm:"column:totalGetMoney"` //(分)
	Money         int32 `gorm:"column:money"`         //(分)
	TodayUseNum   int32 `gorm:"column:todayUseNum"`   //(分)
	UseTime       int64 `gorm:"column:useTime"`
	CashMoney     int64 `gorm:"column:cashMoney"` //现金兑换
	GoldMoney     int64 `gorm:"column:goldMoney"` //元宝兑换
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFeedbackFeeEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFeedbackFeeEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFeedbackFeeEntity) TableName() string {
	return "t_player_feedbackfee"
}
