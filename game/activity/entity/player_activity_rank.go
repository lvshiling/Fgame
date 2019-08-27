package entity

type PlayerActivityRankEntity struct {
	Id           int64  `gorm:"primary_key;column:id"` //Id
	PlayerId     int64  `gorm:"column:playerId"`       //玩家Id
	ActivityType int32  `gorm:"column:activityType"`   //活动类型
	RankMap      string `gorm:"column:rankMap"`        //排行数据
	EndTime      int64  `gorm:"column:endTime"`        //结束时间
	UpdateTime   int64  `gorm:"column:updateTime"`     //更新时间
	CreateTime   int64  `gorm:"column:createTime"`     //创建时间
	DeleteTime   int64  `gorm:"column:deleteTime"`     //删除时间

}

func (e *PlayerActivityRankEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerActivityRankEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerActivityRankEntity) TableName() string {
	return "t_player_activity_rank"
}
