package entity

//玩家档次首充记录数据
type PlayerFirstChargeRecordEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ChargeType int32 `gorm:"column:chargeType"`
	ChargeId   int32 `gorm:"column:chargeId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFirstChargeRecordEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFirstChargeRecordEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFirstChargeRecordEntity) TableName() string {
	return "t_player_first_charge_record"
}
