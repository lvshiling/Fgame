package entity

//玩家圣痕数据
type PlayerShengHenEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerShengHenEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShengHenEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShengHenEntity) TableName() string {
	return "t_player_shenghen"
}
