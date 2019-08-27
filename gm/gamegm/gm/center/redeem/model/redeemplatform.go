package model

type RedeemPlatform struct {
	Id         int64 `gorm:"primary_key;gorm:"column:id"`
	PlatformId int64 `gorm:"column:platformId"`
	RedeemId   int   `gorm:"column:redeemId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (m *RedeemPlatform) TableName() string {
	return "t_redeem_platform"
}
