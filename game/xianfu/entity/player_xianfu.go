package entity

type PlayerXianFuEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	XianfuId   int32 `gorm:"column:xianfuId"`
	XianfuType int32 `gorm:"column:xianfuType"`
	UseTimes   int32 `gorm:"column:useTimes"`
	StartTime  int64 `gorm:"column:startTime"`
	State      int32 `gorm:"column:state"`
	Group      int32 `gorm:"column:group"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerXianFuEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerXianFuEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerXianFuEntity) TableName() string {
	return "t_player_xianfu"
}
