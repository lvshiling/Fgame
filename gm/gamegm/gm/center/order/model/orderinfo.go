package model

type OrderInfo struct {
	Id             int64  `gorm:"primary_key;gorm:"column:id"`
	OrderId        string `gorm:"column:orderId"`
	SdkOrderId     string `gorm:"column:sdkOrderId"`
	Status         int    `gorm:"column:status"`
	SdkType        int    `gorm:"column:sdkType"`
	ServerId       int    `gorm:"column:serverId"`
	UserId         int64  `gorm:"column:userId"`
	PlayerId       int64  `gorm:"column:playerId"`
	ChargeId       int    `gorm:"column:chargeId"`
	Money          int    `gorm:"column:money"`
	ReceivePayTime int64  `gorm:"column:receivePayTime"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
	DevicePlatform int    `gorm:"column:devicePlatform"`
	PlatformUserId string `gorm:"column:platformUserId"`
	PlayerLevel    int    `gorm:"column:playerLevel"`
	PlayerName     string `gorm:"column:playerName"`
	Gold           int    `gorm:"column:gold"`
}

func (m *OrderInfo) TableName() string {
	return "t_order"
}
