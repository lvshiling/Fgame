package model

type PlayerStatEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	BeginTime  int64  `gorm:"column:beginTime"`
	StatType   string `gorm:"column:statType"`
	StatCount  int    `gorm:"column:statCount"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (m *PlayerStatEntity) TableName() string {
	return "t_player_stats"
}
