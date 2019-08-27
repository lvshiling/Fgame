package entity

//玩家活动增长数据
type PlayerActivityAddNumEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	GroupId    int32 `gorm:"column:groupId"`
	AddNum     int32 `gorm:"column:addNum"`
	StartTime  int64 `gorm:"column:startTime"`
	EndTime    int64 `gorm:"column:endTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerActivityAddNumEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerActivityAddNumEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerActivityAddNumEntity) TableName() string {
	return "t_player_activity_add_num"
}
