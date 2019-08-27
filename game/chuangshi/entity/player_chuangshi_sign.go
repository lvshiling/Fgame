package entity

type PlayerChuangShiSignEntity struct {
	Id         int64 `gorm:"column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Status     int32 `gorm:"column:status"` //报名状态：
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerChuangShiSignEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerChuangShiSignEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerChuangShiSignEntity) TableName() string {
	return "t_player_chuangshi_sign"
}
