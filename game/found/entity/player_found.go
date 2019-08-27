package entity

//当天资源记录
type PlayerFoundEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	ResType      int32 `gorm:"column:resType"`
	PlayModeType int32 `gorm:"column:playModeType"`
	JoinTimes    int32 `gorm:"column:joinTimes"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFoundEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFoundEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFoundEntity) TableName() string {
	return "t_player_found"
}
