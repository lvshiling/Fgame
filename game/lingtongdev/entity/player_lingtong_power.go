package entity

//玩家灵童各系统养成战力数据
type PlayerLingTongPowerEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ClassType  int32 `gorm:"column:classType"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerLingTongPowerEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingTongPowerEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingTongPowerEntity) TableName() string {
	return "t_player_lingtong_power"
}
