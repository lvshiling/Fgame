package entity

//玩家buff数据
type PlayerBuffEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	BuffMap    string `gorm:"column:buffMap"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerBuffEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerBuffEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerBuffEntity) TableName() string {
	return "t_player_buff"
}
