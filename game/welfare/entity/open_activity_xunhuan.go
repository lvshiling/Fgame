package entity

//玩家循环活动
type OpenActivityXunHuanEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	ServerId    int32 `gorm:"column:serverId"`
	ArrGroup    int32 `gorm:"column:arrGroup"`
	ActivityDay int32 `gorm:"column:activityDay"`
	StartTime   int64 `gorm:"column:startTime"`
	EndTime     int64 `gorm:"column:endTime"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (e *OpenActivityXunHuanEntity) GetId() int64 {
	return e.Id
}

func (e *OpenActivityXunHuanEntity) TableName() string {
	return "t_open_activity_xun_huan"
}
