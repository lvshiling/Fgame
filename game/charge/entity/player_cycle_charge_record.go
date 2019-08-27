package entity

//玩家每日充值记录数据
type PlayerCycleChargeRecordEntity struct {
	Id              int64 `gorm:"primary_key;column:id"`
	PlayerId        int64 `gorm:"column:playerId"`
	ChargeNum       int64 `gorm:"column:chargeNum"`
	PreDayChargeNum int64 `gorm:"column:preDayChargeNum"`
	UpdateTime      int64 `gorm:"column:updateTime"`
	CreateTime      int64 `gorm:"column:createTime"`
	DeleteTime      int64 `gorm:"column:deleteTime"`
}

func (e *PlayerCycleChargeRecordEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerCycleChargeRecordEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerCycleChargeRecordEntity) TableName() string {
	return "t_player_cycle_charge_record"
}
