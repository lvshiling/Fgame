package model

type CenterPlatformMarryPriceSetInfo struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlatformId int64 `gorm:"column:platformId"`
	KindType   int32 `gorm:"column:kindType"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (m *CenterPlatformMarryPriceSetInfo) TableName() string {
	return "t_platform_marryprice"
}
