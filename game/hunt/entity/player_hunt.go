package entity

//玩家寻宝数据
type PlayerHuntEntity struct {
	Id             int64 `gorm:"primary_key;column:id"`
	PlayerId       int64 `gorm:"column:playerId"`
	HuntType       int32 `gorm:"column:huntType"`
	FreeHuntCount  int32 `gorm:"column:freeHuntCount"`
	TotalHuntCount int32 `gorm:"column:totalHuntCount"`
	LastHuntTime   int64 `gorm:"column:lastHuntTime"`
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (e *PlayerHuntEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerHuntEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerHuntEntity) TableName() string {
	return "t_player_hunt"
}
