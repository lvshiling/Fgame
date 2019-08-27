package model

type PlayerFeedBackFeeRecord struct {
	Id          int64  `gorm:"primary_key;column:id"`
	PlayerId    int64  `gorm:"column:playerId"`
	Money       int32  `gorm:"column:money"`
	Code        string `gorm:"column:code"`
	Status      int32  `gorm:"column:status"`
	Type        int32  `gorm:"column:type"`
	ExpiredTime int64  `gorm:"column:expiredTime"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (m *PlayerFeedBackFeeRecord) TableName() string {
	return "t_player_feedbackfee"
}
