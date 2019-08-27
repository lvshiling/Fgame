package entity

//玩家活跃度数据
type PlayerLivenessEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	Liveness   int64  `gorm:"column:liveness"`
	OpenBoxs   string `gorm:"column:openBoxs"`
	LastTime   int64  `gorm:"column:lastTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerLivenessEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLivenessEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLivenessEntity) TableName() string {
	return "t_player_liveness"
}
