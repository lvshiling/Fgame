package entity

type NewFirstChargeLogEntity struct {
	Id         int64 `gorm:"column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *NewFirstChargeLogEntity) GetId() int64 {
	return e.Id
}

func (e *NewFirstChargeLogEntity) TableName() string {
	return "t_new_first_charge_log"
}
