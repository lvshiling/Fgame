package entity

//玩家boss复活数据
type PlayerBossReliveEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	BossType   int32 `gorm:"column:bossType"`
	ReliveTime int32 `gorm:"column:reliveTime"`

	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerBossReliveEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerBossReliveEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerBossReliveEntity) TableName() string {
	return "t_player_boss_relive"
}
