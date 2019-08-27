package entity

type PlayerWushuangBuchangEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	IsSendEmail int32 `gorm:"column:isSendEmail"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (e *PlayerWushuangBuchangEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerWushuangBuchangEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerWushuangBuchangEntity) TableName() string {
	return "t_player_wushuang_buchang"
}
