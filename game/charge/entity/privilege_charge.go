package entity

//后台充值
type PrivilegeChargeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	Status     int32 `gorm:"column:status"`
	PlayerId   int64 `gorm:"column:playerId"`
	GoldNum    int64 `gorm:"column:goldNum"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PrivilegeChargeEntity) GetId() int64 {
	return e.Id
}

func (e *PrivilegeChargeEntity) TableName() string {
	return "t_privilege_charge"
}
