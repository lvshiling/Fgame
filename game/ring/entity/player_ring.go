package entity

type PlayerRingEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	Typ          int32  `gorm:"column:typ"`
	BindType     int32  `gorm:"column:bindType"`
	ItemId       int32  `gorm:"column:itemId"`
	PropertyData string `gorm:"column:propertyData"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerRingEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerRingEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerRingEntity) TableName() string {
	return "t_player_ring"
}
