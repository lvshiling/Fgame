package entity

//仇人反馈保护
type PlayerFoeProtectEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ExpireTime int64 `gorm:"column:expireTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFoeProtectEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFoeProtectEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFoeProtectEntity) TableName() string {
	return "t_player_foe_protect"
}
