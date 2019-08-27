package entity

// 全部宝宝战力数据
type PlayerBabyPowerEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerBabyPowerEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerBabyPowerEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerBabyPowerEntity) TableName() string {
	return "t_player_baby_power"
}
