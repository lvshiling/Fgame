package entity

//玩家外域BOSS数据
type PlayerOutlandBossEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ZhuoQiNum  int32 `gorm:"column:zhuoqiNum"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerOutlandBossEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerOutlandBossEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerOutlandBossEntity) TableName() string {
	return "t_player_outland_boss"
}
