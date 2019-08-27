package entity

type PlayerChuangShiYuGaoEntity struct {
	Id         int64 `gorm:"column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	IsJoin     int32 `gorm:"column:isJoin"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerChuangShiYuGaoEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerChuangShiYuGaoEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerChuangShiYuGaoEntity) TableName() string {
	return "t_player_chuangshi_yugao"
}
