package model

type PlatformSupportPoolSetInfo struct {
	Id               int64 `gorm:"primary_key;column:id"`
	CenterPlatformId int64 `gorm:"column:centerPlatformId"`
	SupportGold      int32 `gorm:"column:supportGold"`
	SupportRate      int32 `gorm:"column:supportRate"`
	UpdateTime       int64 `gorm:"column:updateTime"`
	CreateTime       int64 `gorm:"column:createTime"`
	DeleteTime       int64 `gorm:"column:deleteTime"`
}

func (m *PlatformSupportPoolSetInfo) TableName() string {
	return "t_platform_supportpool_set"
}
