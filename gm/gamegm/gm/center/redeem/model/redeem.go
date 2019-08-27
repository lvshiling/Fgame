package model

type Redeem struct {
	Id                 int    `gorm:"primary_key;gorm:"column:id"`
	GiftBagName        string `gorm:"column:giftBagName"`
	GiftBagDesc        string `gorm:"column:giftBagDesc"`
	GiftBagContent     string `gorm:"column:giftBagContent"`
	RedeemNum          int    `gorm:"column:redeemNum"`
	RedeemUseNum       int    `gorm:"column:redeemUseNum"`
	RedeemPlayerUseNum int    `gorm:"column:redeemPlayerUseNum"`
	RedeemServerUseNum int    `gorm:"column:redeemServerUseNum"`
	SdkTypes           string `gorm:"column:sdkTypes"`
	SendType           int    `gorm:"column:sendType"`
	StartTime          int64  `gorm:"column:startTime"`
	EndTime            int64  `gorm:"column:endTime"`
	MinPlayerLevel     int    `gorm:"column:minPlayerLevel"`
	MinVipLevel        int    `gorm:"column:minVipLevel"`
	CreateFlag         int    `gorm:"column:createFlag"`
	UpdateTime         int64  `gorm:"column:updateTime"`
	CreateTime         int64  `gorm:"column:createTime"`
	DeleteTime         int64  `gorm:"column:deleteTime"`
}

func (m *Redeem) TableName() string {
	return "t_redeem"
}
