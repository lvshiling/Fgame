package entity

//玩家跨服数据
type PlayerCrossEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	CrossType  int32  `gorm:"column:crossType"`
	CrossArgs  string `gorm:"column:crossArgs"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerCrossEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerCrossEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerCrossEntity) TableName() string {
	return "t_player_cross"
}
