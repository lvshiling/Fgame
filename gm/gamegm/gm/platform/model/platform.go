package model

type PlatformInfo struct {
	PlatformID       int64  `gorm:"primary_key;column:platformId"`
	PlatformName     string `gorm:"column:platformName"`
	CenterPlatformID int64  `gorm:"column:centerPlatformId"`
	ChannelId        int64  `gorm:"column:channelId"`
	UpdateTime       int64  `gorm:"column:updateTime"`
	CreateTime       int64  `gorm:"column:createTime"`
	DeleteTime       int64  `gorm:"column:deleteTime"`
	SdkType          int    `gorm:"column:sdkType"`
	SignKey          string `gorm:"column:signKey"`
}

func (m *PlatformInfo) TableName() string {
	return "t_platform"
}
