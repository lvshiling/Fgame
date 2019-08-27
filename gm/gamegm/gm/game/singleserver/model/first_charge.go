package model

type NewFirstCharge struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	StartTime  int64 `gorm:"column:startTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (m *NewFirstCharge) TableName() string {
	return "t_new_first_charge"
}
