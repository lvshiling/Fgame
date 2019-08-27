package model

type CenterPlatformInfo struct {
	PlatformId int64  `gorm:"primary_key;column:id"`
	Name       string `gorm:"column:name"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
	SkdType    int    `gorm:"column:sdkType"`
}

func (m *CenterPlatformInfo) TableName() string {
	return "t_platform"
}
