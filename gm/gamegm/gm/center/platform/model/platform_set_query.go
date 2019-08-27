package model

type CenterPlatformSetQueryInfo struct {
	Id             int64  `gorm:"primary_key;column:id"`
	PlatformId     int64  `gorm:"column:platformId"`
	PlatformName   string  `gorm:"column:platformName"`
	SettingContent string `gorm:"column:settingContent"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}
