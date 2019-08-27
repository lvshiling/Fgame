package entity

type ActivityEndRecordEntity struct {
	Id           int64 `gorm:"primary_key;column:id"` //Id
	ServerId     int32 `gorm:"column:serverId"`
	ActivityType int32 `gorm:"column:activityType"` //活动类型
	EndTime      int64 `gorm:"column:endTime"`      //活动结束时间
	UpdateTime   int64 `gorm:"column:updateTime"`   //更新时间
	CreateTime   int64 `gorm:"column:createTime"`   //创建时间
	DeleteTime   int64 `gorm:"column:deleteTime"`   //删除时间
}

func (e *ActivityEndRecordEntity) GetId() int64 {
	return e.Id
}

func (e *ActivityEndRecordEntity) TableName() string {
	return "t_activity_end_record"
}
