package entity

//玩家个人BOSS数据
type PlayerMyBossEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	AttendMap  string `gorm:"column:attendMap"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerMyBossEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerMyBossEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerMyBossEntity) TableName() string {
	return "t_player_myboss"
}
