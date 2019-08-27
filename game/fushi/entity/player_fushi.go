package entity

type PlayerFuShiEntity struct {
	Id         int64 `gorm:"column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	FushiLevel int32 `gorm:"column:fushiLevel"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (entity *PlayerFuShiEntity) GetId() int64 {
	return entity.Id
}

func (entity *PlayerFuShiEntity) GetPlayerId() int64 {
	return entity.PlayerId
}

func (entity *PlayerFuShiEntity) TableName() string {
	return "t_player_fushi"
}
