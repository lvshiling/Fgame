package entity

//重置首冲时间
type FirstChargeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	ChargeTime int64 `gorm:"column:chargeTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *FirstChargeEntity) GetId() int64 {
	return e.Id
}

func (e *FirstChargeEntity) TableName() string {
	return "t_first_charge"
}
