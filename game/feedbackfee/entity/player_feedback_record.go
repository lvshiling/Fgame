package entity

//玩家逆付费数据
type PlayerFeedbackRecordEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	PlayerId    int64  `gorm:"column:playerId"`
	Money       int32  `gorm:"column:money"`
	Code        string `gorm:"column:code"`
	Status      int32  `gorm:"column:status"`
	ExpiredTime int64  `gorm:"column:expiredTime"`
	Type        int32  `gorm:"column:type"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (e *PlayerFeedbackRecordEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFeedbackRecordEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFeedbackRecordEntity) TableName() string {
	return "t_player_feedbackfee_record"
}
