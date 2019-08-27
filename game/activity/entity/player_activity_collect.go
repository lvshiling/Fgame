package entity

//玩家采集数据
type PlayerActivityCollectEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	ActivityType int32  `gorm:"column:activityType"`
	CountMap     string `gorm:"column:countMap"`
	EndTime      int64  `gorm:"column:endTime"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerActivityCollectEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerActivityCollectEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerActivityCollectEntity) TableName() string {
	return "t_player_activity_collect"
}
