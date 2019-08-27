package entity

//玩家福利厅
type PlayerOpenActivityEntity struct {
	Id              int64  `gorm:"primary_key;column:id"`
	PlayerId        int64  `gorm:"column:playerId"`
	GroupId         int32  `gorm:"column:groupId"`
	ActivityType    int32  `gorm:"column:activityType"`
	ActivitySubType int32  `gorm:"column:activitySubType"`
	ActivityData    string `gorm:"column:activityData"`
	StartTime       int64  `gorm:"column:startTime"`
	EndTime         int64  `gorm:"column:endTime"`
	UpdateTime      int64  `gorm:"column:updateTime"`
	CreateTime      int64  `gorm:"column:createTime"`
	DeleteTime      int64  `gorm:"column:deleteTime"`
}

func (e *PlayerOpenActivityEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerOpenActivityEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerOpenActivityEntity) TableName() string {
	return "t_player_open_activity"
}
