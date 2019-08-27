package entity

type NewFirstChargeEntity struct {
	Id         int64 `gorm:"column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	StartTime  int64 `gorm:"column:startTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *NewFirstChargeEntity) GetId() int64 {
	return e.Id
}

func (e *NewFirstChargeEntity) TableName() string {
	return "t_new_first_charge"
}
