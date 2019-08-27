package entity

type PlayerActivityEntity struct {
	Id           int64 `gorm:"primary_key;column:id"` //Id
	PlayerId     int64 `gorm:"column:playerId"`       //玩家Id
	ActivityType int32 `gorm:"column:activityType"`   //活动类型
	AttendTimes  int32 `gorm:"column:attendTimes"`    //已参与次数
	UpdateTime   int64 `gorm:"column:updateTime"`     //更新时间
	CreateTime   int64 `gorm:"column:createTime"`     //创建时间
	DeleteTime   int64 `gorm:"column:deleteTime"`     //删除时间

}

func (e *PlayerActivityEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerActivityEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerActivityEntity) TableName() string {
	return "t_player_activity"
}
