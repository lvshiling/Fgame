package model

type PlatformMarrySendLog struct {
	Id          int64  `gorm:"primary_key;column:id"`
	PlatformId  int64  `gorm:"column:platformId"`
	ServerId    int32  `gorm:"column:serverId"`
	SuccessFlag int32  `gorm:"column:successFlag"`
	KindType    int32  `gorm:"column:kindType"`
	FailMsg     string `gorm:"column:failMsg"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (m *PlatformMarrySendLog) TableName() string {
	return "t_marry_set_log"
}
