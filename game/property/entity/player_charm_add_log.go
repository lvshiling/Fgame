package entity

// 魅力增加日志
type PlayerCharmAddLogEntity struct {
	Id         int64 `gorm:"column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SendId     int64 `gorm:"column:sendId"`
	Charm      int32 `gorm:"column:charm"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerCharmAddLogEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerCharmAddLogEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerCharmAddLogEntity) TableName() string {
	return "t_player_charm_add_log"
}
