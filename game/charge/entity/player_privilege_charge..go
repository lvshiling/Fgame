package entity

//玩家充值数据
type PlayerPrivilegeChargeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ChargeType int32 `gorm:"column:chargeType"`
	ChargeId   int32 `gorm:"column:chargeId"`
	ChargeNum  int32 `gorm:"column:chargeNum"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerPrivilegeChargeEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerPrivilegeChargeEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerPrivilegeChargeEntity) TableName() string {
	return "t_player_privilege_charge"
}
