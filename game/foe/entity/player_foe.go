package entity

//仇人列表
type PlayerFoeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	AttackId   int64 `gorm:"column:foeId"`
	KillTime   int64 `gorm:"column:killTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFoeEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFoeEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFoeEntity) TableName() string {
	return "t_player_foe"
}
