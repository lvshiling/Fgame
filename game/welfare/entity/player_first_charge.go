package entity

//玩家首充
type PlayerFirstChargeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	IsReceive  bool  `gorm:"column:isReceive"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFirstChargeEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFirstChargeEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFirstChargeEntity) TableName() string {
	return "t_player_first_charge"
}
