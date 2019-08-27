package model

type RedeemCode struct {
	Id         int64  `gorm:"primary_key;gorm:"column:id"`
	RedeemCode string `gorm:"column:redeemCode"`
	RedeemId   int    `gorm:"column:redeemId"`
	UseNum     int    `gorm:"column:useNum"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (m *RedeemCode) TableName() string {
	return "t_redeem_code"
}
