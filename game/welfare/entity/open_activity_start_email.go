package entity

//开服活动开启邮件记录数据
type OpenActivityStartEmailEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	EndTime    int64 `gorm:"column:endTime"`
	GroupId    int32 `gorm:"column:groupId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *OpenActivityStartEmailEntity) GetId() int64 {
	return e.Id
}

func (e *OpenActivityStartEmailEntity) TableName() string {
	return "t_open_activity_start_mail"
}
