package model

type NoticeInfo struct {
	Id               int64  `gorm:"primary_key;column:id"`
	ChannelId        int    `gorm:"column:channelId"`
	PlatformId       int    `gorm:"column:platformId"`
	ServerId         int    `gorm:"column:serverId"`
	Content          string `gorm:"column:content"`
	BeginTime        int64  `gorm:"column:beginTime"`
	EndTime          int64  `gorm:"column:endTime"`
	IntervalTime     int64  `gorm:"column:intervalTime"`
	UpdateTime       int64  `gorm:"column:updateTime"`
	CreateTime       int64  `gorm:"column:createTime"`
	DeleteTime       int64  `gorm:"column:deleteTime"`
	SuccessFlag      int    `gorm:"column:successFlag"`
	ErrorMsg         string `gorm:"column:errorMsg"`
	ServerName       string `gorm:"column:serverName"`
	CenterPlatformId int64  `gorm:"column:centerPlatformId"`
	UserName         string `gorm:"column:userName"`
}

func (m *NoticeInfo) TableName() string {
	return "t_notice"
}
