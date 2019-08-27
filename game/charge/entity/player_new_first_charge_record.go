package entity

type PlayerNewFirstChargeRecordEntity struct {
	Id         int64  `gorm:"column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	Record     string `gorm:"column:record"`
	StartTime  int64  `gorm:"column:startTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerNewFirstChargeRecordEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerNewFirstChargeRecordEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerNewFirstChargeRecordEntity) TableName() string {
	return "t_player_new_first_charge_record"
}
