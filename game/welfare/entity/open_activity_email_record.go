package entity

//开服排行榜邮件奖励记录数据
type OpenActivityEmailRecordEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	EndTime    int64 `gorm:"column:endTime"`
	GroupId    int32 `gorm:"column:groupId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *OpenActivityEmailRecordEntity) GetId() int64 {
	return e.Id
}

func (e *OpenActivityEmailRecordEntity) TableName() string {
	return "t_open_activity_email_record"
}
