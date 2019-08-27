package entity

//玩家活动抽奖
type PlayerActivityNumRecordEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	GroupId    int32 `gorm:"column:groupId"`
	Times      int32 `gorm:"column:times"`
	StartTime  int64 `gorm:"column:startTime"`
	EndTime    int64 `gorm:"column:endTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerActivityNumRecordEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerActivityNumRecordEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerActivityNumRecordEntity) TableName() string {
	return "t_player_activity_num_record"
}
